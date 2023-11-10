package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type ReplaceOneOptions struct {
	Filter   any
	Unscoped bool
	Options  []*options.ReplaceOptions
}

func NewReplaceOneOption() *ReplaceOneOptions {
	return new(ReplaceOneOptions)
}

type ReplaceOneOption interface {
	ApplyReplaceOne(*ReplaceOneOptions)
}
