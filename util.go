package mongox

import (
	"errors"
	"reflect"

	"github.com/lilith44/easy"
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
		return easy.Underscore(element.Name())
	case reflect.Struct:
		return easy.Underscore(value.Type().Name())
	default:
		panic("getCollectionName: needs a struct or slice, or a pointer to them")
	}

	return ""
}
func parseFindResult(err error) (bool, error) {
	if err == nil {
		return true, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return false, err
}
