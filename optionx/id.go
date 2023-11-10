package optionx

import "go.mongodb.org/mongo-driver/bson"

type Id struct {
	id any
}

func WithId(id any) *Id {
	return &Id{
		id: id,
	}
}

func (i *Id) ApplyFindOne(options *FindOneOptions) {
	options.Filter = bson.M{"_id": i.id}
}

func (i *Id) ApplyReplaceOne(options *ReplaceOneOptions) {
	options.Filter = bson.M{"_id": i.id}
}

func (i *Id) ApplyUpdateOne(options *UpdateOneOptions) {
	options.Filter = bson.M{"_id": i.id}
}
