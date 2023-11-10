package fields

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Field struct {
	Id       primitive.ObjectID `bson:"_id"`
	UpdateAt int64              `mongox:"update_at" bson:"update_at"`
	DeleteAt *int64             `mongox:"delete_at" bson:"delete_at,omitempty"`
}
