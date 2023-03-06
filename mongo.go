package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	database *mongo.Database
}

func New(c Config) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.URI).SetAuth(c.Auth))
	if err != nil {
		return nil, err
	}

	return &Mongo{database: client.Database(c.Scheme)}, nil
}

func (m *Mongo) Database() *mongo.Database {
	return m.database
}

func (m *Mongo) InsertOne(
	ctx context.Context,
	collection string, doc any, options ...*options.InsertOneOptions,
) (*mongo.InsertOneResult, error) {
	return m.database.Collection(collection).InsertOne(ctx, doc, options...)
}

func (m *Mongo) InsertMany(
	ctx context.Context,
	collection string, docs []any, options ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	if len(docs) == 0 {
		return nil, nil
	}

	return m.database.Collection(collection).InsertMany(ctx, docs, options...)
}

func (m *Mongo) UpdateOne(
	ctx context.Context,
	collection string, filter any, update any, options ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return m.database.Collection(collection).UpdateOne(ctx, filter, update, options...)
}

func (m *Mongo) UpdateById(
	ctx context.Context,
	collection string, id any, update any, options ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return m.database.Collection(collection).UpdateByID(ctx, id, update, options...)
}

func (m *Mongo) UpdateMany(
	ctx context.Context,
	collection string, id any, update any, options ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return m.database.Collection(collection).UpdateMany(ctx, id, update, options...)
}

func (m *Mongo) ReplaceOne(
	ctx context.Context,
	collection string, filter any, replacement any, options ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	return m.database.Collection(collection).ReplaceOne(ctx, filter, replacement, options...)
}

func (m *Mongo) DeleteOne(
	ctx context.Context,
	collection string, filter any, options ...*options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	return m.database.Collection(collection).DeleteOne(ctx, filter, options...)
}

func (m *Mongo) BulkWrite(
	ctx context.Context,
	collection string, models []mongo.WriteModel, options ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	return m.database.Collection(collection).BulkWrite(ctx, models, options...)
}

func (m *Mongo) CountDocument(ctx context.Context, collection string, filer any) (int64, error) {
	return m.database.Collection(collection).CountDocuments(ctx, filer)
}

func (m *Mongo) EstimatedDocumentCount(
	ctx context.Context,
	collection string, options ...*options.EstimatedDocumentCountOptions,
) (int64, error) {
	return m.database.Collection(collection).EstimatedDocumentCount(ctx, options...)
}

func (m *Mongo) Find(
	ctx context.Context,
	collection string, filter any, options ...*options.FindOptions,
) (*mongo.Cursor, error) {
	return m.database.Collection(collection).Find(ctx, filter, options...)
}

func (m *Mongo) FindOne(
	ctx context.Context,
	collection string, filter any, options ...*options.FindOneOptions,
) *mongo.SingleResult {
	return m.database.Collection(collection).FindOne(ctx, filter, options...)
}

func (m *Mongo) FindOneAndDelete(
	ctx context.Context,
	collection string, filter any, options ...*options.FindOneAndDeleteOptions,
) *mongo.SingleResult {
	return m.database.Collection(collection).FindOneAndDelete(ctx, filter, options...)
}

func (m *Mongo) FindOneAndUpdate(
	ctx context.Context,
	collection string, filter any, update any, options ...*options.FindOneAndUpdateOptions,
) *mongo.SingleResult {
	return m.database.Collection(collection).FindOneAndUpdate(ctx, filter, update, options...)
}

func (m *Mongo) FindOneAndReplace(
	ctx context.Context,
	collection string, filter any, replacement any, options ...*options.FindOneAndReplaceOptions,
) *mongo.SingleResult {
	return m.database.Collection(collection).FindOneAndReplace(ctx, filter, replacement, options...)
}

func (m *Mongo) Aggregate(
	ctx context.Context,
	collection string, pipeline any, options ...*options.AggregateOptions,
) (*mongo.Cursor, error) {
	return m.database.Collection(collection).Aggregate(ctx, pipeline, options...)
}

func (m *Mongo) Transaction(ctx context.Context, f func(session mongo.SessionContext) error) error {
	return m.database.Client().UseSession(ctx, func(session mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err
		}

		defer session.EndSession(context.Background())

		if err = f(session); err != nil {
			_ = session.AbortTransaction(context.Background())
			return err
		}
		return session.CommitTransaction(context.Background())
	})
}
