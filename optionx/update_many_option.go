package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type updateManyOptions struct {
	options []*options.UpdateOptions
}

func WithUpdateManyOptions(options ...*options.UpdateOptions) UpdateManyOption {
	return &updateManyOptions{
		options: options,
	}
}

func (uo *updateManyOptions) ApplyUpdateMany(options *UpdateManyOptions) {
	options.Options = append(options.Options, uo.options...)
}
