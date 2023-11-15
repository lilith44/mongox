package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type AggregateOptions struct {
	Unscoped bool
	Options  []*options.AggregateOptions
}

func NewAggregateOption() *AggregateOptions {
	return new(AggregateOptions)
}

type AggregateOption interface {
	ApplyAggregate(*AggregateOptions)
}
