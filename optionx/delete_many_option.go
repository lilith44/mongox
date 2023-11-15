package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type deleteManyOptions struct {
	options []*options.DeleteOptions
}

func WithDeleteManyOptions(options ...*options.DeleteOptions) DeleteManyOption {
	return &deleteManyOptions{
		options: options,
	}
}

func (do *deleteManyOptions) ApplyDeleteMany(options *DeleteManyOptions) {
	options.Options = append(options.Options, do.options...)
}
