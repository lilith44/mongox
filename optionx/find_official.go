package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type findOfficial struct {
	official []*options.FindOptions
}

func WithFindOfficial(official ...*options.FindOptions) FindOption {
	return &findOfficial{
		official: official,
	}
}

func (fo *findOfficial) ApplyFind(options *FindOptions) {
	options.Options = append(options.Options, fo.official...)
}
