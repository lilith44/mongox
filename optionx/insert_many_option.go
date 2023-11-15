package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type insertManyOptions struct {
	options []*options.InsertManyOptions
}

func WithInsertManyOptions(options ...*options.InsertManyOptions) InsertManyOption {
	return &insertManyOptions{
		options: options,
	}
}

func (io *insertManyOptions) ApplyInsertMany(options *InsertManyOptions) {
	options.Options = append(options.Options, io.options...)
}
