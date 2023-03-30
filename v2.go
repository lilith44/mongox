package mongox

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type V2 struct {
	database *mongo.Database
	logger   *zap.SugaredLogger

	collection string
	options    any
	filter     any
	pipeline   any
	models     []mongo.WriteModel
	limit      *int64
	skip       *int64
	sort       bson.D
}

func (v *V2) Collection(collection string) *V2 {
	v.collection = collection
	return v
}

func (v *V2) Options(options any) *V2 {
	value := reflect.ValueOf(options)
	switch value.Kind() {
	case reflect.Pointer:
		opt := reflect.MakeSlice(reflect.SliceOf(value.Type()), 0, 1)
		opt = reflect.Append(opt, value)
		v.options = opt.Interface()
	case reflect.Slice:
		v.options = options
	}
	return v
}

func (v *V2) Filter(filter any) *V2 {
	v.filter = filter
	return v
}

func (v *V2) ObjectId(id any) *V2 {
	v.filter = bson.M{"_id": id}
	return v
}

func (v *V2) Pipeline(pipeline any) *V2 {
	v.pipeline = pipeline
	return v
}

func (v *V2) Models(models []mongo.WriteModel) *V2 {
	v.models = models
	return v
}

func (v *V2) Limit(limit int64) *V2 {
	v.limit = &limit
	return v
}

func (v *V2) Skip(skip int64) *V2 {
	v.skip = &skip
	return v
}

func (v *V2) Sort(sorts ...bson.E) *V2 {
	v.sort = append(v.sort, sorts...)
	return v
}

func (v *V2) Asc(field string) *V2 {
	v.sort = append(v.sort, bson.E{Key: field, Value: 1})
	return v
}

func (v *V2) Desc(field string) *V2 {
	v.sort = append(v.sort, bson.E{Key: field, Value: -1})
	return v
}

// InsertOne 未指定collection时，根据bean获取collectionName
func (v *V2) InsertOne(ctx context.Context, bean any) (any, error) {
	if v.collection == "" {
		v.collection = getCollectionName(bean)
	}

	var _options []*options.InsertOneOptions
	if v.options != nil {
		_options = v.options.([]*options.InsertOneOptions)
	}

	v.logger.Infof("Add document %+v into collection %s", bean, v.collection)

	r, err := v.database.Collection(v.collection).InsertOne(ctx, bean, _options...)
	if err != nil {
		return nil, err
	}
	return r.InsertedID, nil
}

// InsertMany 未指定collection时，根据beans获取collectionName
func (v *V2) InsertMany(ctx context.Context, beans ...any) (any, error) {
	if len(beans) == 0 {
		return nil, nil
	}

	if v.collection == "" {
		v.collection = getCollectionName(beans)
	}

	var _options []*options.InsertManyOptions
	if v.options != nil {
		_options = v.options.([]*options.InsertManyOptions)
	}

	v.logger.Infof("Add multi documents %+v into collection %s", beans, v.collection)

	r, err := v.database.Collection(v.collection).InsertMany(ctx, beans, _options...)
	if err != nil {
		return nil, err
	}
	return r.InsertedIDs, nil
}

func (v *V2) UpdateOne(ctx context.Context, update any) (*mongo.UpdateResult, error) {
	var _options []*options.UpdateOptions
	if v.options != nil {
		_options = v.options.([]*options.UpdateOptions)
	}

	v.logger.Infof("Update one document in collection %s, filter: %s, update: %s", v.collection, v.filter, update)

	return v.database.Collection(v.collection).UpdateOne(ctx, v.filter, update, _options...)
}

func (v *V2) UpdateById(ctx context.Context, id any, update any) (*mongo.UpdateResult, error) {
	var _options []*options.UpdateOptions
	if v.options != nil {
		_options = v.options.([]*options.UpdateOptions)
	}

	v.logger.Infof("Update one document in collection %s, id: %s, update: %s", v.collection, id, update)

	return v.database.Collection(v.collection).UpdateByID(ctx, id, update, _options...)
}

func (v *V2) UpdateMany(ctx context.Context, update any) (*mongo.UpdateResult, error) {
	var _options []*options.UpdateOptions
	if v.options != nil {
		_options = v.options.([]*options.UpdateOptions)
	}

	v.logger.Infof("Update multi documents in collection %s, filter: %s, update: %s", v.collection, v.filter, update)

	return v.database.Collection(v.collection).UpdateMany(ctx, v.filter, update, _options...)
}

// ReplaceOne 未指定collection时，根据replacement获取collectionName
func (v *V2) ReplaceOne(ctx context.Context, replacement any) (*mongo.UpdateResult, error) {
	if v.collection == "" {
		v.collection = getCollectionName(replacement)
	}

	var _options []*options.UpdateOptions
	if v.options != nil {
		_options = v.options.([]*options.UpdateOptions)
	}

	v.logger.Infof("Replace one document in collection %s, filter: %s, replacement: %+v", v.collection, v.filter, replacement)

	return v.database.Collection(v.collection).UpdateOne(ctx, v.filter, replacement, _options...)
}

func (v *V2) DeleteOne(ctx context.Context) (int64, error) {
	var _options []*options.DeleteOptions
	if v.options != nil {
		_options = v.options.([]*options.DeleteOptions)
	}

	v.logger.Infof("Delete one document from collection %s, filter: %s", v.collection, v.filter)

	r, err := v.database.Collection(v.collection).DeleteOne(ctx, v.filter, _options...)
	if err != nil {
		return 0, err
	}
	return r.DeletedCount, nil
}

func (v *V2) DeleteMany(ctx context.Context) (int64, error) {
	if v.filter == nil {
		panic("Delete multi documents without filter")
	}

	var _options []*options.DeleteOptions
	if v.options != nil {
		_options = v.options.([]*options.DeleteOptions)
	}

	v.logger.Infof("Delete multi documents from collection %s, filter: %s", v.collection, v.filter)

	r, err := v.database.Collection(v.collection).DeleteMany(ctx, v.filter, _options...)
	if err != nil {
		return 0, err
	}
	return r.DeletedCount, nil
}

func (v *V2) BulkWrite(ctx context.Context) (*mongo.BulkWriteResult, error) {
	var _options []*options.BulkWriteOptions
	if v.options != nil {
		_options = v.options.([]*options.BulkWriteOptions)
	}

	v.logger.Infof("BulkWrite documents into collection %s, models: %+v", v.collection, v.models)

	return v.database.Collection(v.collection).BulkWrite(ctx, v.models, _options...)
}

func (v *V2) CountDocument(ctx context.Context) (int64, error) {
	var _options []*options.CountOptions
	if v.options != nil {
		_options = v.options.([]*options.CountOptions)
	}

	v.logger.Infof("Count documents in collection %s, filter: %s", v.collection, v.filter)

	return v.database.Collection(v.collection).CountDocuments(ctx, v.filter, _options...)
}

func (v *V2) EstimatedDocumentCount(ctx context.Context) (int64, error) {
	var _options []*options.EstimatedDocumentCountOptions
	if v.options != nil {
		_options = v.options.([]*options.EstimatedDocumentCountOptions)
	}

	v.logger.Infof("Estimated document count in collection %s", v.collection)

	return v.database.Collection(v.collection).EstimatedDocumentCount(ctx, _options...)
}

// Find 未指定collection时，根据beansPtr获取collectionName。beansPtr需要是指向切片的指针
func (v *V2) Find(ctx context.Context, beansPtr any) error {
	if v.collection == "" {
		v.collection = getCollectionName(beansPtr)
	}

	var _options []*options.FindOptions
	if v.options != nil {
		_options = v.options.([]*options.FindOptions)
	}

	if v.limit != nil {
		_options = append(_options, &options.FindOptions{Limit: v.limit})
	}
	if v.skip != nil {
		_options = append(_options, &options.FindOptions{Skip: v.skip})
	}
	if v.sort != nil {
		_options = append(_options, &options.FindOptions{Sort: v.sort})
	}

	v.logger.Infof("Find documents in collection %s, filter: %s", v.collection, v.filter)

	r, err := v.database.Collection(v.collection).Find(ctx, v.filter, _options...)
	if err != nil {
		return err
	}
	return r.All(ctx, beansPtr)
}

// FindOne 未指定collection时，根据beanPtr获取collectionName。beanPtr需要是指针类型
func (v *V2) FindOne(ctx context.Context, beanPtr any) (bool, error) {
	if v.collection == "" {
		v.collection = getCollectionName(beanPtr)
	}

	var _options []*options.FindOneOptions
	if v.options != nil {
		_options = v.options.([]*options.FindOneOptions)
	}

	if v.skip != nil {
		_options = append(_options, &options.FindOneOptions{Skip: v.skip})
	}
	if v.sort != nil {
		_options = append(_options, &options.FindOneOptions{Sort: v.sort})
	}

	v.logger.Infof("Find one document in collection %s, filter: %s", v.collection, v.filter)

	err := v.database.Collection(v.collection).FindOne(ctx, v.filter, _options...).Decode(beanPtr)
	return parseFindOneResult(err)
}

// FindOneAndDelete 未指定collection时，根据beanPtr获取collectionName。beanPtr需要是指针类型
func (v *V2) FindOneAndDelete(ctx context.Context, beanPtr any) (bool, error) {
	if v.collection == "" {
		v.collection = getCollectionName(beanPtr)
	}

	var _options []*options.FindOneAndDeleteOptions
	if v.options != nil {
		_options = v.options.([]*options.FindOneAndDeleteOptions)
	}

	if v.sort != nil {
		_options = append(_options, &options.FindOneAndDeleteOptions{Sort: v.sort})
	}

	v.logger.Infof("Find one document and delete it in collection %s, filter: %s", v.collection, v.filter)

	err := v.database.Collection(v.collection).FindOneAndDelete(ctx, v.filter, _options...).Decode(beanPtr)
	return parseFindOneResult(err)
}

// FindOneAndUpdate 未指定collection时，根据beanPtr获取collectionName。beanPtr需要是指针类型
func (v *V2) FindOneAndUpdate(ctx context.Context, beanPtr any, update any) (bool, error) {
	var _options []*options.FindOneAndUpdateOptions
	if v.options != nil {
		_options = v.options.([]*options.FindOneAndUpdateOptions)
	}

	if v.sort != nil {
		_options = append(_options, &options.FindOneAndUpdateOptions{Sort: v.sort})
	}

	v.logger.Infof("Find one document and update it in collection %s, filter: %s, update: %+v", v.collection, v.filter, update)

	err := v.database.Collection(v.collection).FindOneAndUpdate(ctx, v.filter, update, _options...).Decode(beanPtr)
	return parseFindOneResult(err)
}

// FindOneAndReplace 未指定collection时，根据beanPtr获取collectionName。beanPtr需要是指针类型
func (v *V2) FindOneAndReplace(ctx context.Context, beanPtr any, replacement any) (bool, error) {
	if v.collection == "" {
		v.collection = getCollectionName(beanPtr)
	}

	var _options []*options.FindOneAndReplaceOptions
	if v.options != nil {
		_options = v.options.([]*options.FindOneAndReplaceOptions)
	}

	if v.sort != nil {
		_options = append(_options, &options.FindOneAndReplaceOptions{Sort: v.sort})
	}

	v.logger.Infof("Find one document and replace it in collection %s, filter: %s, replacement: %+v", v.collection, v.filter, replacement)

	err := v.database.Collection(v.collection).FindOneAndReplace(ctx, v.filter, replacement, _options...).Decode(beanPtr)
	return parseFindOneResult(err)
}

func (v *V2) Aggregate(ctx context.Context) (*mongo.Cursor, error) {
	var _options []*options.AggregateOptions
	if v.options != nil {
		_options = v.options.([]*options.AggregateOptions)
	}

	v.logger.Infof("Aggregate in collection %s, pipeline: %s", v.collection, v.pipeline)

	return v.database.Collection(v.collection).Aggregate(ctx, v.pipeline, _options...)
}

func (v *V2) Transaction(ctx context.Context, f func(session mongo.SessionContext) error) error {
	return v.database.Client().UseSession(ctx, func(session mongo.SessionContext) error {
		v.logger.Infof("Begin transaction []")

		err := session.StartTransaction()
		if err != nil {
			return err
		}

		defer func() {
			session.EndSession(context.Background())
			if err != nil {
				v.logger.Errorf("Error accurred in transaction: %s", err)
			}
		}()

		if err = f(session); err != nil {
			_ = session.AbortTransaction(context.Background())
			v.logger.Errorf("Abort []")
			return err
		}

		v.logger.Infof("Commit []")
		return session.CommitTransaction(context.Background())
	})
}
