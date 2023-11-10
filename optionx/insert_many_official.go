package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type insertManyOfficial struct {
	official []*options.InsertManyOptions
}

func WithInsertManyOfficial(official ...*options.InsertManyOptions) InsertManyOption {
	return &insertManyOfficial{
		official: official,
	}
}

func (io *insertManyOfficial) ApplyInsertMany(options *InsertManyOptions) {
	options.Options = append(options.Options, io.official...)
}
