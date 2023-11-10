package mongox

import (
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type field struct {
	indexes []int
	name    string
	typ     reflect.Type
}

type collection struct {
	name string

	updateAtFields []*field
	deleteAtFields []*field
}

func (c *collection) parseFields(value reflect.Value, tag string, index []int) {
	value = reflect.Indirect(value)
	if value.Kind() != reflect.Struct {
		return
	}

	t := value.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous && strings.Contains(f.Tag.Get("bson"), "inline") {
			c.parseFields(value.Field(i), tag, append(index, i))
			continue
		}

		if f.Tag.Get(tag) == "update_at" {
			name := strings.Split(f.Tag.Get("bson"), ",")[0]
			if name == "" {
				name = f.Name
			}

			c.updateAtFields = append(c.updateAtFields, &field{
				indexes: append(index, i),
				name:    name,
				typ:     value.Field(i).Type(),
			})
		}

		if f.Tag.Get(tag) == "delete_at" {
			name := strings.Split(f.Tag.Get("bson"), ",")[0]
			if name == "" {
				name = f.Name
			}

			typ := value.Field(i).Type()
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}
			c.deleteAtFields = append(c.deleteAtFields, &field{
				indexes: append(index, i),
				name:    name,
				typ:     typ,
			})
		}
	}
}

func (c *collection) setUpdateAt(bean any) {
	value := reflect.Indirect(reflect.ValueOf(bean))
	for _, f := range c.updateAtFields {
		for _, i := range f.indexes {
			value = value.Field(i)
		}

		switch f.typ.Kind() {
		case reflect.Int64:
			value.Set(reflect.ValueOf(time.Now().Unix()))
		case reflect.Struct:
			value.Set(reflect.ValueOf(time.Now()))
		}
	}
}

func (c *collection) getFilterForQuery(filter any) any {
	for _, f := range c.deleteAtFields {
		filter = setDeleteAtForQuery(filter, f.name, false)
	}
	return filter
}

func (c *collection) getUpdateForUpdate(update any) any {
	for _, f := range c.updateAtFields {
		var value any
		switch f.typ.Kind() {
		case reflect.Int64:
			value = time.Now().Unix()
		case reflect.Struct:
			value = time.Now()
		}

		update = setUpdateAtForWrite(update, f.name, value)
	}
	return update
}

func (c *collection) getUpdateForDelete() bson.M {
	update := make(bson.M)
	for _, f := range c.deleteAtFields {
		var value any
		switch f.typ.Kind() {
		case reflect.Int64:
			value = time.Now().Unix()
		case reflect.Struct:
			value = time.Now()
		}

		update[f.name] = value
	}
	return bson.M{Set: update}
}

func (c *collection) getPipelineForAggregate(pipeline any) any {
	for _, f := range c.deleteAtFields {
		pipeline = setDeleteAtForAggregate(pipeline, f.name, false)
	}
	return pipeline
}
