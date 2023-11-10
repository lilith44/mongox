package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type replaceOneOfficial struct {
	official []*options.ReplaceOptions
}

func WithReplaceOneOfficial(official ...*options.ReplaceOptions) ReplaceOneOption {
	return &replaceOneOfficial{
		official: official,
	}
}

func (ro *replaceOneOfficial) ApplyReplaceOne(options *ReplaceOneOptions) {
	options.Options = append(options.Options, ro.official...)
}
