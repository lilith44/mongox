package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type deleteManyOfficial struct {
	official []*options.DeleteOptions
}

func WithDeleteManyOfficial(official ...*options.DeleteOptions) DeleteManyOption {
	return &deleteManyOfficial{
		official: official,
	}
}

func (do *deleteManyOfficial) ApplyDeleteMany(options *DeleteManyOptions) {
	options.Options = append(options.Options, do.official...)
}
