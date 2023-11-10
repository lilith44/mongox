package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type FindOptions struct {
	Collection string
	Filter     any
	Unscoped   bool
	Options    []*options.FindOptions
}

func NewFindOption() *FindOptions {
	return new(FindOptions)
}

type FindOption interface {
	ApplyFind(*FindOptions)
}
