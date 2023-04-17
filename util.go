package mongox

import (
	"errors"
	"reflect"

	"github.com/lilith44/easy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollectionName(s any) string {
	value := reflect.Indirect(reflect.ValueOf(s))

	switch value.Kind() {
	case reflect.Slice:
		element := value.Type().Elem()
		if element.Kind() == reflect.Pointer {
			element = element.Elem()
		}
		if element.Kind() == reflect.Interface {
			if value.Len() == 0 {
				panic("slice []any must not be empty")
			}
			return getCollectionName(value.Index(0).Interface())
		}
		return easy.Underscore(element.Name())
	case reflect.Struct:
		return easy.Underscore(value.Type().Name())
	default:
		panic("getCollectionName: needs a struct or slice, or a pointer to them")
	}

	return ""
}

func getUpdateFieldsByStruct(bean any, mustFields ...string) bson.D {
	value := reflect.Indirect(reflect.ValueOf(bean))
	if value.Kind() != reflect.Struct {
		panic("getUpdateFieldsByStruct: needs a struct, or a pointer to it")
	}

	var (
		filter    bson.D
		f         func(value reflect.Value, prefix string)
		mustField = make(map[string]bool)
	)

	for _, field := range mustFields {
		mustField[field] = true
	}

	f = func(value reflect.Value, prefix string) {
		typ := value.Type()
		for i := 0; i < typ.NumField(); i++ {
			t := typ.Field(i)
			if !t.IsExported() {
				continue
			}

			key := t.Tag.Get("bson")
			if key == "" {
				key = t.Name
			}
			key = prefix + key

			if len(mustField) != 0 && !mustField[key] {
				continue
			}

			v := value.Field(i)
			if mustField[key] {
				filter = append(filter, bson.E{Key: key, Value: v.Interface()})
				continue
			}

			if v.Kind() == reflect.Pointer {
				if v.IsNil() {
					continue
				}
				v = v.Elem()
			}

			if v.IsZero() {
				continue
			}

			if v.Kind() == reflect.Struct {
				f(v, key+".")

				continue
			}

			filter = append(filter, bson.E{Key: key, Value: v.Interface()})
		}
	}

	f(value, "")

	return filter
}

func parseFindOneResult(err error) (bool, error) {
	if err == nil {
		return true, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return false, err
}
