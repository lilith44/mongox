package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type findOptions struct {
	options []*options.FindOptions
}

func WithFindOptions(options ...*options.FindOptions) FindOption {
	return &findOptions{
		options: options,
	}
}

func (fo *findOptions) ApplyFind(options *FindOptions) {
	options.Options = append(options.Options, fo.options...)
}
