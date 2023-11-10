package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type CountOptions struct {
	Filter   any
	Unscoped bool
	Options  []*options.CountOptions
}

func NewCountOption() *CountOptions {
	return new(CountOptions)
}

type CountOption interface {
	ApplyCount(*CountOptions)
}
