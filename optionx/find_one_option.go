package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type findOneOptions struct {
	options []*options.FindOneOptions
}

func WithFindOneOptions(options ...*options.FindOneOptions) FindOneOption {
	return &findOneOptions{
		options: options,
	}
}

func (fo *findOneOptions) ApplyFindOne(options *FindOneOptions) {
	options.Options = append(options.Options, fo.options...)
}
