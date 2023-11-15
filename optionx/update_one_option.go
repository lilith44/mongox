package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type updateOneOptions struct {
	options []*options.UpdateOptions
}

func WithUpdateOneOptions(options ...*options.UpdateOptions) UpdateOneOption {
	return &updateOneOptions{
		options: options,
	}
}

func (uo *updateOneOptions) ApplyUpdateOne(options *UpdateOneOptions) {
	options.Options = append(options.Options, uo.options...)
}
