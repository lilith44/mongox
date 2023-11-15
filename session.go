package mongox

import (
	"context"

	"github.com/lilith44/mongox/optionx"
	"github.com/lilith44/mongox/timer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Mongo) InsertOne(
	ctx context.Context,
	bean any,
	options ...optionx.InsertOneOption,
) (result *mongo.InsertOneResult, err error) {
	o := optionx.NewInsertOneOption()
	for _, option := range options {
		option.ApplyInsertOne(o)
	}

	collectionName := getCollectionName(bean)
	if o.Collection == "" {
		o.Collection = collectionName
	}

	m.collections[collectionName].setUpdateAt(bean)

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("InsertOne"),
		timer.WithCollection(o.Collection),
		timer.WithBean(bean),
	)
	defer t.End(err)

	result, err = m.database.Collection(o.Collection).InsertOne(ctx, bean, o.Options...)
	return
}

func (m *Mongo) InsertMany(
	ctx context.Context,
	beans []any,
	options ...optionx.InsertManyOption,
) (result *mongo.InsertManyResult, err error) {
	o := optionx.NewInsertManyOption()
	for _, option := range options {
		option.ApplyInsertMany(o)
	}

	collectionName := getCollectionName(beans)
	if o.Collection == "" {
		o.Collection = collectionName
	}

	for _, bean := range beans {
		m.collections[collectionName].setUpdateAt(bean)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("InsertMany"),
		timer.WithCollection(o.Collection),
		timer.WithBean(beans),
	)
	defer t.End(err)

	result, err = m.database.Collection(o.Collection).InsertMany(ctx, beans, o.Options...)
	return
}

func (m *Mongo) UpdateOne(
	ctx context.Context,
	collection string,
	update any,
	options ...optionx.UpdateOneOption,
) (result *mongo.UpdateResult, err error) {
	o := optionx.NewUpdateOneOption()
	for _, option := range options {
		option.ApplyUpdateOne(o)
	}
	if o.Filter == nil {
		o.Filter = make(bson.M)
	}
	if !o.Unscoped {
		o.Filter = m.collections[collection].getFilterForQuery(o.Filter)
	}
	update = m.collections[collection].getUpdateForUpdate(update)

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("UpdateOne"),
		timer.WithCollection(collection),
		timer.WithFilter(o.Filter),
		timer.WithUpdate(update),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).UpdateOne(ctx, o.Filter, update, o.Options...)
	return
}

func (m *Mongo) UpdateMany(
	ctx context.Context,
	collection string,
	update any,
	options ...optionx.UpdateManyOption,
) (result *mongo.UpdateResult, err error) {
	o := optionx.NewUpdateManyOption()
	for _, option := range options {
		option.ApplyUpdateMany(o)
	}
	if o.Filter == nil {
		o.Filter = make(bson.M)
	}
	if !o.Unscoped {
		o.Filter = m.collections[collection].getFilterForQuery(o.Filter)
	}
	update = m.collections[collection].getUpdateForUpdate(update)

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("UpdateMany"),
		timer.WithCollection(collection),
		timer.WithFilter(o.Filter),
		timer.WithUpdate(update),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).UpdateMany(ctx, o.Filter, update, o.Options...)
	return
}

func (m *Mongo) ReplaceOne(
	ctx context.Context,
	collection string,
	replacement any,
	options ...optionx.ReplaceOneOption,
) (result *mongo.UpdateResult, err error) {
	o := optionx.NewReplaceOneOption()
	for _, option := range options {
		option.ApplyReplaceOne(o)
	}
	if o.Filter == nil {
		o.Filter = make(bson.M)
	}
	if !o.Unscoped {
		o.Filter = m.collections[collection].getFilterForQuery(o.Filter)
	}
	m.collections[collection].setUpdateAt(replacement)

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("ReplaceOne"),
		timer.WithCollection(collection),
		timer.WithFilter(o.Filter),
		timer.WithUpdate(replacement),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).ReplaceOne(ctx, o.Filter, replacement, o.Options...)
	return
}

func (m *Mongo) DeleteOne(
	ctx context.Context,
	collection string,
	filter any,
	options ...optionx.DeleteOneOption,
) (result *mongo.DeleteResult, err error) {
	o := optionx.NewDeleteOneOption()
	for _, option := range options {
		option.ApplyDeleteOne(o)
	}

	if !o.Unscoped && len(m.collections[collection].deleteAtFields) != 0 {
		var updateResult *mongo.UpdateResult
		if updateResult, err = m.UpdateOne(ctx, collection, m.collections[collection].getUpdateForDelete(), optionx.WithFilter(filter)); err != nil {
			return
		}

		result = &mongo.DeleteResult{DeletedCount: updateResult.ModifiedCount}
		return
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("DeleteOne"),
		timer.WithCollection(collection),
		timer.WithFilter(filter),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).DeleteOne(ctx, filter, o.Options...)
	return
}

func (m *Mongo) DeleteMany(
	ctx context.Context,
	collection string,
	filter any,
	options ...optionx.DeleteManyOption,
) (result *mongo.DeleteResult, err error) {
	o := optionx.NewDeleteManyOption()
	for _, option := range options {
		option.ApplyDeleteMany(o)
	}

	if !o.Unscoped && len(m.collections[collection].deleteAtFields) != 0 {
		var updateResult *mongo.UpdateResult
		if updateResult, err = m.UpdateOne(ctx, collection, m.collections[collection].getUpdateForDelete(), optionx.WithFilter(filter)); err != nil {
			return
		}

		result = &mongo.DeleteResult{DeletedCount: updateResult.ModifiedCount}
		return
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("DeleteMany"),
		timer.WithCollection(collection),
		timer.WithFilter(filter),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).DeleteMany(ctx, filter, o.Options...)
	return
}

func (m *Mongo) BulkWrite(
	ctx context.Context,
	collection string,
	models []mongo.WriteModel,
	options ...optionx.BulkWriteOption,
) (result *mongo.BulkWriteResult, err error) {
	o := optionx.NewBulkWriteOption()
	for _, option := range options {
		option.ApplyBulkWrite(o)
	}

	for _, model := range models {
		m.collections[collection].setUpdateAt(model)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("BulkWrite"),
		timer.WithCollection(collection),
		timer.WithModel(models),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).BulkWrite(ctx, models, o.Options...)
	return
}

func (m *Mongo) CountDocument(
	ctx context.Context,
	collection string,
	options ...optionx.CountOption,
) (result int64, err error) {
	o := optionx.NewCountOption()
	for _, option := range options {
		option.ApplyCount(o)
	}
	if o.Filter == nil {
		o.Filter = bson.M{}
	}
	if !o.Unscoped {
		o.Filter = m.collections[collection].getFilterForQuery(o.Filter)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("CountDocument"),
		timer.WithCollection(collection),
		timer.WithFilter(o.Filter),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).CountDocuments(ctx, o.Filter, o.Options...)
	return
}

func (m *Mongo) EstimatedDocumentCount(
	ctx context.Context,
	collection string,
	options ...optionx.EstimatedDocumentCountOption,
) (result int64, err error) {
	o := optionx.NewEstimatedDocumentCountOption()
	for _, option := range options {
		option.ApplyEstimatedDocumentCount(o)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("EstimatedDocumentCount"),
		timer.WithCollection(collection),
	)
	defer t.End(err)

	result, err = m.database.Collection(collection).EstimatedDocumentCount(ctx, o.Options...)
	return
}

func (m *Mongo) Find(ctx context.Context, beansPtr any, options ...optionx.FindOption) (err error) {
	o := optionx.NewFindOption()
	for _, option := range options {
		option.ApplyFind(o)
	}

	collectionName := getCollectionName(beansPtr)
	if o.Filter == nil {
		o.Filter = make(bson.M)
	}
	if !o.Unscoped {
		o.Filter = m.collections[collectionName].getFilterForQuery(o.Filter)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("Find"),
		timer.WithCollection(collectionName),
		timer.WithFilter(o.Filter),
	)
	defer t.End(err)

	var cursor *mongo.Cursor
	if cursor, err = m.database.Collection(collectionName).Find(ctx, o.Filter, o.Options...); err != nil {
		return
	}

	err = cursor.All(ctx, beansPtr)
	return
}

func (m *Mongo) FindOne(ctx context.Context, beanPtr any, options ...optionx.FindOneOption) (exist bool, err error) {
	o := optionx.NewFindOneOption()
	for _, option := range options {
		option.ApplyFindOne(o)
	}

	collectionName := getCollectionName(beanPtr)
	if o.Filter == nil {
		o.Filter = make(bson.M)
	}
	if !o.Unscoped {
		o.Filter = m.collections[collectionName].getFilterForQuery(o.Filter)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("FindOne"),
		timer.WithCollection(collectionName),
		timer.WithFilter(o.Filter),
	)
	defer t.End(err)

	err = m.database.Collection(collectionName).FindOne(ctx, o.Filter, o.Options...).Decode(beanPtr)
	exist, err = parseFindOneResult(err)
	return
}

func (m *Mongo) Aggregate(
	ctx context.Context,
	collection string,
	pipeline any,
	options ...optionx.AggregateOption,
) (cursor *mongo.Cursor, err error) {
	o := optionx.NewAggregateOption()
	for _, option := range options {
		option.ApplyAggregate(o)
	}
	if !o.Unscoped {
		pipeline = m.collections[collection].getPipelineForAggregate(pipeline)
	}

	t := timer.New(
		timer.WithLogger(m.logger),
		timer.WithMethod("Aggregate"),
		timer.WithCollection(collection),
		timer.WithPipeline(pipeline),
	)
	defer t.End(err)

	cursor, err = m.database.Collection(collection).Aggregate(ctx, pipeline, o.Options...)
	return
}

func (m *Mongo) Transaction(ctx context.Context, f func(session mongo.SessionContext) error) error {
	return m.database.Client().UseSession(ctx, func(session mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			m.logger.Errorf("start transaction failed: %s", err)

			return err
		}

		m.logger.Infof("transaction start...")

		defer func() {
			session.EndSession(ctx)
			if err != nil {
				m.logger.Errorf("Error occurred in transaction: %s", err)
			}
		}()

		if err = f(session); err != nil {
			_ = session.AbortTransaction(ctx)
			m.logger.Errorf("transaction aborted...")
			return err
		}

		m.logger.Infof("transaction ended...")
		return session.CommitTransaction(ctx)
	})
}
