package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type EstimatedDocumentCountOptions struct {
	Options []*options.EstimatedDocumentCountOptions
}

func NewEstimatedDocumentCountOption() *EstimatedDocumentCountOptions {
	return new(EstimatedDocumentCountOptions)
}

type EstimatedDocumentCountOption interface {
	ApplyEstimatedDocumentCount(*EstimatedDocumentCountOptions)
}
