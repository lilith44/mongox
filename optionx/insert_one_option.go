package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type insertOneOptions struct {
	options []*options.InsertOneOptions
}

func WithInsertOneOptions(options ...*options.InsertOneOptions) InsertOneOption {
	return &insertOneOptions{
		options: options,
	}
}

func (io *insertOneOptions) ApplyInsertOne(options *InsertOneOptions) {
	options.Options = append(options.Options, io.options...)
}
