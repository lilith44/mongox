package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type DeleteOneOptions struct {
	Unscoped bool
	Options  []*options.DeleteOptions
}

func NewDeleteOneOption() *DeleteOneOptions {
	return new(DeleteOneOptions)
}

type DeleteOneOption interface {
	ApplyDeleteOne(*DeleteOneOptions)
}
