package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type findOneOfficial struct {
	official []*options.FindOneOptions
}

func WithFindOneOfficial(official ...*options.FindOneOptions) FindOneOption {
	return &findOneOfficial{
		official: official,
	}
}

func (fo *findOneOfficial) ApplyFindOne(options *FindOneOptions) {
	options.Options = append(options.Options, fo.official...)
}
