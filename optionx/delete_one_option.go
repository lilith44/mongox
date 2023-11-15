package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type deleteOneOptions struct {
	options []*options.DeleteOptions
}

func WithDeleteOneOptions(options ...*options.DeleteOptions) DeleteOneOption {
	return &deleteOneOptions{
		options: options,
	}
}

func (do *deleteOneOptions) ApplyDeleteOne(options *DeleteOneOptions) {
	options.Options = append(options.Options, do.options...)
}
