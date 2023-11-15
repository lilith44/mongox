package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type estimatedDocumentCountOptions struct {
	options []*options.EstimatedDocumentCountOptions
}

func WithEstimatedDocumentCountOptions(options ...*options.EstimatedDocumentCountOptions) EstimatedDocumentCountOption {
	return &estimatedDocumentCountOptions{
		options: options,
	}
}

func (eo *estimatedDocumentCountOptions) ApplyEstimatedDocumentCount(options *EstimatedDocumentCountOptions) {
	options.Options = append(options.Options, eo.options...)
}
