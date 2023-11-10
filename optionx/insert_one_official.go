package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type insertOneOfficial struct {
	official []*options.InsertOneOptions
}

func WithInsertOneOfficial(official ...*options.InsertOneOptions) InsertOneOption {
	return &insertOneOfficial{
		official: official,
	}
}

func (io *insertOneOfficial) ApplyInsertOne(options *InsertOneOptions) {
	options.Options = append(options.Options, io.official...)
}
