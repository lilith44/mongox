package mongox

import (
	"reflect"
)

func (m *Mongo) RegisterModels(models ...any) {
	for _, model := range models {
		c := &collection{
			name: getCollectionName(model),
		}

		c.parseFields(reflect.ValueOf(model), "mongox", nil)
		m.collections[c.name] = c
	}
}
