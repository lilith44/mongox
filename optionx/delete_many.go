package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type DeleteManyOptions struct {
	Unscoped bool
	Options  []*options.DeleteOptions
}

func NewDeleteManyOption() *DeleteManyOptions {
	return new(DeleteManyOptions)
}

type DeleteManyOption interface {
	ApplyDeleteMany(*DeleteManyOptions)
}
