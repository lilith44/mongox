package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type InsertManyOptions struct {
	Options []*options.InsertManyOptions
}

func NewInsertManyOption() *InsertManyOptions {
	return new(InsertManyOptions)
}

type InsertManyOption interface {
	ApplyInsertMany(*InsertManyOptions)
}
