package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type UpdateManyOptions struct {
	Filter   any
	Unscoped bool
	Options  []*options.UpdateOptions
}

func NewUpdateManyOption() *UpdateManyOptions {
	return new(UpdateManyOptions)
}

type UpdateManyOption interface {
	ApplyUpdateMany(*UpdateManyOptions)
}
