package mongox

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type T1 struct {
	Id    string        `bson:"id"`
	Name  string        `bson:"name"`
	Chan  chan struct{} `bson:"chan"`
	Slice []string
	Map   map[string]any `bson:"map"`
	T2    *NestedT2      `bson:"t2"`
}

type NestedT2 struct {
	Id   string  `bson:"_id"`
	Name *string `bson:"name"`
	Bool bool    `bson:"bool"`
}

func TestGetUpdateFieldsByStruct(t *testing.T) {
	type Case struct {
		Input    T1
		MustCols []string
		Output   bson.D
	}

	c := make(chan struct{})
	s := []string{"1", "2", "3"}
	m := map[string]any{"1": 2, "2": 3}
	t2 := &NestedT2{
		Id:   "3",
		Name: nil,
		Bool: false,
	}

	suits := []Case{
		{
			Input: T1{
				Id:    "1",
				Name:  "",
				Chan:  c,
				Slice: s,
				Map:   m,
				T2: &NestedT2{
					Id:   "3",
					Name: nil,
					Bool: false,
				},
			},
			Output: bson.D{
				{Key: "id", Value: "1"},
				{Key: "chan", Value: c},
				{Key: "Slice", Value: s},
				{Key: "map", Value: m},
				{Key: "t2._id", Value: "3"},
			},
		},
		{
			Input: T1{
				Id:    "1",
				Name:  "",
				Chan:  c,
				Slice: s,
				Map:   m,
				T2:    t2,
			},
			MustCols: []string{"id", "t2"},
			Output: bson.D{
				{Key: "id", Value: "1"},
				{Key: "t2", Value: t2},
			},
		},
	}

	for _, suit := range suits {
		if !reflect.DeepEqual(getUpdateFieldsByStruct(suit.Input, suit.MustCols...), suit.Output) {
			t.Fail()
		}
	}
}
