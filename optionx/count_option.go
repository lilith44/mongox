package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type countOptions struct {
	options []*options.CountOptions
}

func WithCountOptions(options ...*options.CountOptions) CountOption {
	return &countOptions{
		options: options,
	}
}

func (co *countOptions) ApplyCount(options *CountOptions) {
	options.Options = append(options.Options, co.options...)
}
