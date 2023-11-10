package mongox

import (
	"cmp"
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

func parseFindOneResult(err error) (bool, error) {
	if err == nil {
		return true, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return false, err
}

func set(filter any, field string, value any) any {
	if filter == nil {
		return filter
	}

	switch f := filter.(type) {
	case bson.M:
		f[field] = value
		return f
	case *bson.M:
		(*f)[field] = value
		return f
	case bson.D:
		f = append(f, bson.E{Key: field, Value: value})
		return f
	case *bson.D:
		*f = append(*f, bson.E{Key: field, Value: value})
		return f
	}
	return filter
}

func setUpdateAtForWrite(update any, field string, value any) any {
	if update == nil {
		return update
	}

	switch u := update.(type) {
	case bson.M:
		if _, ok := u[Set]; ok {
			u[Set] = set(u[Set], field, value)
		}
		return u
	case *bson.M:
		if _, ok := (*u)[Set]; ok {
			(*u)[Set] = set((*u)[Set], field, value)
		}
		return u
	case bson.D:
		for i := range u {
			if u[i].Key == Set {
				u[i].Value = set(u[i].Value, field, value)
				return u
			}
		}
	case *bson.D:
		for i := range *u {
			if (*u)[i].Key == Set {
				(*u)[i].Value = set((*u)[i].Value, field, value)
				return u
			}
		}
	}
	return update
}

func setDeleteAtForQuery(filter any, field string, exists bool) any {
	if filter == nil {
		return filter
	}

	sub := bson.M{Exists: exists}
	switch f := filter.(type) {
	case bson.M:
		f[field] = sub
		return f
	case *bson.M:
		(*f)[field] = sub
		return f
	case bson.D:
		f = append(f, bson.E{Key: field, Value: sub})
		return f
	case *bson.D:
		*f = append(*f, bson.E{Key: field, Value: sub})
		return f
	}
	return filter
}

func setDeleteAtForAggregate(pipeline any, field string, exists bool) any {
	switch p := pipeline.(type) {
	case bson.A:
		if len(p) > 0 {
			switch pp := p[0].(type) {
			case bson.M:
				if _, ok := pp[Match]; ok {
					pp[Match] = set(pp[Match], field, bson.M{Exists: exists})
					return p
				}
			case bson.D:
				if len(pp) > 0 && pp[0].Key == Match {
					pp[0].Value = set(pp[0].Value, field, bson.M{Exists: exists})
					return p
				}
			}
		}
		p = append(bson.A{bson.M{Match: bson.M{field: bson.M{Exists: exists}}}}, p...)
		return p
	case []bson.D:
		if len(p) > 0 {
			pp := p[0]
			if len(pp) > 0 && pp[0].Key == Match {
				pp[0].Value = set(pp[0].Value, field, bson.M{Exists: exists})
				return p
			}
		}
		p = append([]bson.D{{{Key: Match, Value: bson.M{field: bson.M{Exists: exists}}}}}, p...)
		return p
	case mongo.Pipeline:
		if len(p) > 0 {
			pp := p[0]
			if len(pp) > 0 && pp[0].Key == Match {
				pp[0].Value = set(pp[0].Value, field, bson.M{Exists: exists})
				return p
			}
		}
		p = append([]bson.D{{{Key: Match, Value: bson.M{field: bson.M{Exists: exists}}}}}, p...)
		return p
	case *bson.A:
		if len(*p) > 0 {
			switch pp := (*p)[0].(type) {
			case bson.M:
				if _, ok := pp[Match]; ok {
					pp[Match] = set(pp[Match], field, bson.M{Exists: exists})
					return p
				}
			case bson.D:
				if len(pp) > 0 && pp[0].Key == Match {
					pp[0].Value = set(pp[0].Value, field, bson.M{Exists: exists})
					return p
				}
			}
		}
		*p = append(bson.A{bson.M{Match: bson.M{field: bson.M{Exists: exists}}}}, *p...)
		return p
	case *[]bson.D:
		if len(*p) > 0 {
			pp := (*p)[0]
			if len(pp) > 0 && pp[0].Key == Match {
				pp[0].Value = set(pp[0].Value, field, bson.M{Exists: exists})
				return p
			}
		}
		*p = append([]bson.D{{{Key: Match, Value: bson.M{field: bson.M{Exists: exists}}}}}, *p...)
		return p
	case *mongo.Pipeline:
		if len(*p) > 0 {
			pp := (*p)[0]
			if len(pp) > 0 && pp[0].Key == Match {
				pp[0].Value = set(pp[0].Value, field, bson.M{Exists: exists})
				return p
			}
		}
		*p = append([]bson.D{{{Key: Match, Value: bson.M{field: bson.M{Exists: exists}}}}}, *p...)
		return p

	}
	return pipeline
}

func diffIndexes(existIndexes []*index, indexes []mongo.IndexModel) ([]string, []mongo.IndexModel) {
	var toDelete []string
	for _, existIndex := range existIndexes {
		var exist bool
		for _, index := range indexes {
			if indexEqual(existIndex.Key, index.Keys.(bson.D)) {
				exist = true
				break
			}
		}
		if !exist && existIndex.Key[0].Key != "_id" {
			toDelete = append(toDelete, existIndex.Name)
		}
	}

	var toCreate []mongo.IndexModel
	for _, index := range indexes {
		var exist bool
		for _, existIndex := range existIndexes {
			if indexEqual(existIndex.Key, index.Keys.(bson.D)) {
				exist = true
				break
			}
		}
		if !exist {
			toCreate = append(toCreate, index)
		}
	}
	return toDelete, toCreate
}

func indexEqual(x, y bson.D) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i].Key != y[i].Key || !integerEqual(x[i].Value, y[i].Value) {
			return false
		}
	}
	return true
}

func integerEqual(x, y any) bool {
	vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
	if vx.CanInt() && vy.CanInt() {
		return cmp.Compare(vx.Int(), vy.Int()) == 0
	}
	if vx.CanUint() && vy.CanUint() {
		return cmp.Compare(vx.Uint(), vy.Uint()) == 0
	}
	return false
}
