package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type bulkWriteOfficial struct {
	official []*options.BulkWriteOptions
}

func WithBulkWriteOfficial(official ...*options.BulkWriteOptions) BulkWriteOption {
	return &bulkWriteOfficial{
		official: official,
	}
}

func (bo *bulkWriteOfficial) ApplyBulkWrite(options *BulkWriteOptions) {
	options.Options = append(options.Options, bo.official...)
}
