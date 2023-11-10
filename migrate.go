package mongox

import (
	"context"

	"mongox/timer"

	"github.com/lilith44/easy/slicex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MigrateOption struct {
	Collection        string
	CollectionOptions []*options.CreateCollectionOptions

	Indexes      []mongo.IndexModel
	IndexOptions []*options.CreateIndexesOptions
}

// Mirage creates collection(s) if not exist(s), creates new indexes and removes missed indexes.
func (m *Mongo) Mirage(ctx context.Context, options ...*MigrateOption) error {
	if len(options) == 0 {
		return nil
	}

	collections, err := m.database.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return err
	}

	collectionMap := slicex.ToMap(collections)
	for _, option := range options {
		if _, ok := collectionMap[option.Collection]; !ok {
			t := timer.New(
				timer.WithLogger(m.logger),
				timer.WithMethod("Create"),
				timer.WithCollection(option.Collection),
			)
			err = m.database.CreateCollection(ctx, option.Collection, option.CollectionOptions...)
			t.End(err)
			if err != nil {
				return err
			}
		}

		if err = m.migrateIndexes(ctx, option.Collection, option.Indexes, option.IndexOptions); err != nil {
			return err
		}
	}
	return nil
}

type index struct {
	Name string `bson:"name"`
	Key  bson.D `bson:"key"`
}

func (m *Mongo) migrateIndexes(ctx context.Context, collection string, indexes []mongo.IndexModel, options []*options.CreateIndexesOptions) error {
	cursor, err := m.database.Collection(collection).Indexes().List(ctx)
	if err != nil {
		return err
	}

	var existIndexes []*index
	if err = cursor.All(ctx, &existIndexes); err != nil {
		return err
	}

	toDelete, toCreate := diffIndexes(existIndexes, indexes)
	for _, name := range toDelete {
		t := timer.New(
			timer.WithLogger(m.logger),
			timer.WithMethod("DropOneIndex"),
			timer.WithCollection(collection),
			timer.WithBean(name),
		)

		_, err = m.database.Collection(collection).Indexes().DropOne(ctx, name)
		t.End(err)
		if err != nil {
			return err
		}
	}
	if len(toCreate) != 0 {
		t := timer.New(
			timer.WithLogger(m.logger),
			timer.WithMethod("CreateManyIndex"),
			timer.WithCollection(collection),
			timer.WithBean(toCreate),
		)
		_, err = m.database.Collection(collection).Indexes().CreateMany(ctx, toCreate, options...)
		t.End(err)
		if err != nil {
			return err
		}
	}
	return nil
}
