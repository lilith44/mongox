package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type replaceOneOptions struct {
	options []*options.ReplaceOptions
}

func WithReplaceOneOptions(options ...*options.ReplaceOptions) ReplaceOneOption {
	return &replaceOneOptions{
		options: options,
	}
}

func (ro *replaceOneOptions) ApplyReplaceOne(options *ReplaceOneOptions) {
	options.Options = append(options.Options, ro.options...)
}
