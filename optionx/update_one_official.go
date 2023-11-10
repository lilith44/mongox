package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type updateOneOfficial struct {
	official []*options.UpdateOptions
}

func WithUpdateOneOfficial(official ...*options.UpdateOptions) UpdateOneOption {
	return &updateOneOfficial{
		official: official,
	}
}

func (uo *updateOneOfficial) ApplyUpdateOne(options *UpdateOneOptions) {
	options.Options = append(options.Options, uo.official...)
}
