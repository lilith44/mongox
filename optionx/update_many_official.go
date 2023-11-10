package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type updateManyOfficial struct {
	official []*options.UpdateOptions
}

func WithUpdateManyOfficial(official ...*options.UpdateOptions) UpdateManyOption {
	return &updateManyOfficial{
		official: official,
	}
}

func (uo *updateManyOfficial) ApplyUpdateMany(options *UpdateManyOptions) {
	options.Options = append(options.Options, uo.official...)
}
