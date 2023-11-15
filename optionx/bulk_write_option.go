package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type bulkWriteOptions struct {
	options []*options.BulkWriteOptions
}

func WithBulkWriteOptions(options ...*options.BulkWriteOptions) BulkWriteOption {
	return &bulkWriteOptions{
		options: options,
	}
}

func (bo *bulkWriteOptions) ApplyBulkWrite(options *BulkWriteOptions) {
	options.Options = append(options.Options, bo.options...)
}
