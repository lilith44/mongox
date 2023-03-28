package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Mongo struct {
	database *mongo.Database

	logger *zap.SugaredLogger
}

type SyncCollectionOption struct {
	Collection string
	Options    []*options.CreateCollectionOptions
}

func New(c Config, logger *zap.SugaredLogger) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.URI).SetAuth(c.Auth))
	if err != nil {
		return nil, err
	}

	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &Mongo{
		database: client.Database(c.Scheme),
		logger:   logger,
	}, nil
}

func (m *Mongo) Database() *mongo.Database {
	return m.database
}

func (m *Mongo) Sync(ctx context.Context, options ...*SyncCollectionOption) error {
	if len(options) == 0 {
		return nil
	}

	collections, err := m.database.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return err
	}

	collectionMap := make(map[string]bool, len(collections))
	for _, collection := range collections {
		collectionMap[collection] = true
	}

	for _, option := range options {
		if collectionMap[option.Collection] {
			continue
		}

		if err = m.database.CreateCollection(ctx, option.Collection, option.Options...); err != nil {
			return err
		}
	}
	return nil
}

func (m *Mongo) InsertOne(
	ctx context.Context,
	bean any, options ...*options.InsertOneOptions,
) (any, error) {
	m.logger.Infof("Add document %+v into collection %s", bean, getCollectionName(bean))

	r, err := m.database.Collection(getCollectionName(bean)).InsertOne(ctx, bean, options...)
	if err != nil {
		return nil, err
	}
	return r.InsertedID, nil
}

func (m *Mongo) InsertMany(
	ctx context.Context,
	beans []any, options ...*options.InsertManyOptions,
) ([]any, error) {
	if len(beans) == 0 {
		return nil, nil
	}

	m.logger.Infof("Add multi documents %+v into collection %s", beans, getCollectionName(beans))

	r, err := m.database.Collection(getCollectionName(beans)).InsertMany(ctx, beans, options...)
	if err != nil {
		return nil, err
	}
	return r.InsertedIDs, nil
}

func (m *Mongo) UpdateOne(
	ctx context.Context,
	collection string, filter any, update any, options ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	m.logger.Infof("Update one document in collection %s, filter: %s, update: %s", collection, filter, update)

	return m.database.Collection(collection).UpdateOne(ctx, filter, update, options...)
}

func (m *Mongo) UpdateById(
	ctx context.Context,
	collection string, id any, update any, options ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	m.logger.Infof("Update one document in collection %s, id: %s, update: %s", collection, id, update)

	return m.database.Collection(collection).UpdateByID(ctx, id, update, options...)
}

func (m *Mongo) UpdateMany(
	ctx context.Context,
	collection string, filter any, update any, options ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	m.logger.Infof("Update multi documents in collection %s, filter: %s, update: %s", collection, filter, update)

	return m.database.Collection(collection).UpdateMany(ctx, filter, update, options...)
}

func (m *Mongo) ReplaceOne(
	ctx context.Context,
	collection string, filter any, replacement any, options ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	m.logger.Infof("Replace one document in collection %s, filter: %s, replacement: %s", collection, filter, replacement)

	return m.database.Collection(collection).ReplaceOne(ctx, filter, replacement, options...)
}

func (m *Mongo) DeleteOne(
	ctx context.Context,
	collection string, filter any, options ...*options.DeleteOptions,
) (int64, error) {
	m.logger.Infof("Delete one documents in collection %s, filter: %s", collection, filter)

	r, err := m.database.Collection(collection).DeleteOne(ctx, filter, options...)
	if err != nil {
		return 0, err
	}
	return r.DeletedCount, nil
}

func (m *Mongo) BulkWrite(
	ctx context.Context,
	collection string, models []mongo.WriteModel, options ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	m.logger.Infof("BulkWrite documents in collection %s, models: %+v", collection, models)

	return m.database.Collection(collection).BulkWrite(ctx, models, options...)
}

func (m *Mongo) CountDocument(ctx context.Context, collection string, filer any, options ...*options.CountOptions) (int64, error) {
	m.logger.Infof("Count documents in collection %s, filter: %s", collection, filer)

	return m.database.Collection(collection).CountDocuments(ctx, filer, options...)
}

func (m *Mongo) EstimatedDocumentCount(
	ctx context.Context,
	collection string, options ...*options.EstimatedDocumentCountOptions,
) (int64, error) {
	m.logger.Infof("Estimated document count in collection %s", collection)

	return m.database.Collection(collection).EstimatedDocumentCount(ctx, options...)
}

func (m *Mongo) Find(
	ctx context.Context,
	beansPtr any, filter any, options ...*options.FindOptions,
) error {
	m.logger.Infof("Find documents in collection %s, filter: %s", getCollectionName(beansPtr), filter)

	r, err := m.database.Collection(getCollectionName(beansPtr)).Find(ctx, filter, options...)
	if err != nil {
		return err
	}
	return r.All(ctx, beansPtr)
}

func (m *Mongo) FindOne(
	ctx context.Context,
	beanPtr any, filter any, options ...*options.FindOneOptions,
) (bool, error) {
	m.logger.Infof("Find one document in collection %s, filter: %s", getCollectionName(beanPtr), filter)

	err := m.database.Collection(getCollectionName(beanPtr)).FindOne(ctx, filter, options...).Decode(beanPtr)
	return parseFindResult(err)
}

func (m *Mongo) FindOneAndDelete(
	ctx context.Context,
	beanPtr any, filter any, options ...*options.FindOneAndDeleteOptions,
) (bool, error) {
	m.logger.Infof("Find one document in collection %s and delete, filter: %s", getCollectionName(beanPtr), filter)

	err := m.database.Collection(getCollectionName(beanPtr)).FindOneAndDelete(ctx, filter, options...).Decode(beanPtr)
	return parseFindResult(err)
}

func (m *Mongo) FindOneAndUpdate(
	ctx context.Context,
	beanPtr any, filter any, update any, options ...*options.FindOneAndUpdateOptions,
) (bool, error) {
	m.logger.Infof("Find one document in collection %s and update, filter: %s, update: %s", getCollectionName(beanPtr), filter, update)

	err := m.database.Collection(getCollectionName(beanPtr)).FindOneAndUpdate(ctx, filter, update, options...).Decode(beanPtr)
	return parseFindResult(err)
}

func (m *Mongo) FindOneAndReplace(
	ctx context.Context,
	beanPtr any, filter any, replacement any, options ...*options.FindOneAndReplaceOptions,
) (bool, error) {
	m.logger.Infof("Find one document in collection %s and replace, filter: %s, replacement: %s", getCollectionName(beanPtr), filter, replacement)

	err := m.database.Collection(getCollectionName(beanPtr)).FindOneAndReplace(ctx, filter, replacement, options...).Decode(beanPtr)
	return parseFindResult(err)
}

func (m *Mongo) Aggregate(
	ctx context.Context,
	collection string, pipeline any, options ...*options.AggregateOptions,
) (*mongo.Cursor, error) {
	m.logger.Infof("Aggregate in collection %s, pipeline: %s", collection, pipeline)

	return m.database.Collection(collection).Aggregate(ctx, pipeline, options...)
}

func (m *Mongo) Transaction(ctx context.Context, f func(session mongo.SessionContext) error) error {
	return m.database.Client().UseSession(ctx, func(session mongo.SessionContext) error {
		m.logger.Infof("Begin transaction []")

		err := session.StartTransaction()
		if err != nil {
			return err
		}

		defer func() {
			session.EndSession(context.Background())
			if err != nil {
				m.logger.Errorf("Error accurred in transaction: %s", err)
			}
		}()

		if err = f(session); err != nil {
			_ = session.AbortTransaction(context.Background())
			m.logger.Errorf("Abort []")
			return err
		}

		m.logger.Infof("Commit []")
		return session.CommitTransaction(context.Background())
	})
}
