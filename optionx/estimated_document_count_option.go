package optionx

import "go.mongodb.org/mongo-driver/mongo/options"

type estimatedDocumentCountOfficial struct {
	official []*options.EstimatedDocumentCountOptions
}

func WithEstimatedDocumentCountOfficial(official ...*options.EstimatedDocumentCountOptions) EstimatedDocumentCountOption {
	return &estimatedDocumentCountOfficial{
		official: official,
	}
}

func (eo *estimatedDocumentCountOfficial) ApplyEstimatedDocumentCount(options *EstimatedDocumentCountOptions) {
	options.Options = append(options.Options, eo.official...)
}
