package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type UpdateOneOptions struct {
	Filter   any
	Unscoped bool
	Options  []*options.UpdateOptions
}

func NewUpdateOneOption() *UpdateOneOptions {
	return new(UpdateOneOptions)
}

type UpdateOneOption interface {
	ApplyUpdateOne(*UpdateOneOptions)
}
