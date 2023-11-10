package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type InsertOneOptions struct {
	Collection string
	Options    []*options.InsertOneOptions
}

func NewInsertOneOption() *InsertOneOptions {
	return new(InsertOneOptions)
}

type InsertOneOption interface {
	ApplyInsertOne(*InsertOneOptions)
}
