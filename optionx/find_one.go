package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type FindOneOptions struct {
	Filter   any
	Unscoped bool
	Options  []*options.FindOneOptions
}

func NewFindOneOption() *FindOneOptions {
	return new(FindOneOptions)
}

type FindOneOption interface {
	ApplyFindOne(*FindOneOptions)
}
