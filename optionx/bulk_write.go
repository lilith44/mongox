package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type BulkWriteOptions struct {
	Options []*options.BulkWriteOptions
}

func NewBulkWriteOption() *BulkWriteOptions {
	return new(BulkWriteOptions)
}

type BulkWriteOption interface {
	ApplyBulkWrite(*BulkWriteOptions)
}
