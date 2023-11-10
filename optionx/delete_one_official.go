package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type deleteOneOfficial struct {
	official []*options.DeleteOptions
}

func WithDeleteOneOfficial(official ...*options.DeleteOptions) DeleteOneOption {
	return &deleteOneOfficial{
		official: official,
	}
}

func (do *deleteOneOfficial) ApplyDeleteOne(options *DeleteOneOptions) {
	options.Options = append(options.Options, do.official...)
}
