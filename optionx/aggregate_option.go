package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type aggregateOptions struct {
	options []*options.AggregateOptions
}

func WithAggregateOptions(options ...*options.AggregateOptions) AggregateOption {
	return &aggregateOptions{
		options: options,
	}
}

func (ao *aggregateOptions) ApplyAggregate(options *AggregateOptions) {
	options.Options = ao.options
}
