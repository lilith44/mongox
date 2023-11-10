package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type countOfficial struct {
	official []*options.CountOptions
}

func WithCountOfficial(official ...*options.CountOptions) CountOption {
	return &countOfficial{
		official: official,
	}
}

func (co *countOfficial) ApplyCount(options *CountOptions) {
	options.Options = append(options.Options, co.official...)
}
