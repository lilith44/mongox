package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type aggregateOfficial struct {
	official []*options.AggregateOptions
}

func WithAggregateOfficial(official ...*options.AggregateOptions) AggregateOption {
	return &aggregateOfficial{
		official: official,
	}
}

func (ao *aggregateOfficial) ApplyAggregate(options *AggregateOptions) {
	options.Options = ao.official
}
