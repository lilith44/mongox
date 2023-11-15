# mongox
基于官方包go.mongodb.org/mongo-driver进行简单封装，主要有以下特性：

+ 支持update_at字段
+ 支持软删除
+ 支持动态迁移collection以及其index
+ 支持日志打印
+ 减少了部分官方包的方法的参数数量

需要go版本在1.21.0以上。

# 用法
## 创建一个mongox实例
``` go
logger, _ := zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel))
m, err := mongox.New(
	mongox.Config{
		Scheme: "x",
		URI:    "mongodb://localhost:27017,localhost:27018,localhost:27019",
		Auth: mongox.Auth{
			AuthMechanism: "SCRAM-SHA-256",
			AuthSource:    "admin",
			Username:      "root",
			Password:      "123456",
		},
	},
	logger.Sugar(),
)
if err != nil {
	return err
}

```

## 定义你的model
创建一个user model：
``` go
type User struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
}

```

若需要添加一个随更新操作而更新为当前时间的字段，那么可以：
``` go
type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
	UpdateAt int64              `mongox:"update_at" bson:"update_at"`
}

```
其中包含mongox:"update_at"这个tag的字段，会在UpdateOne，UpdateMany，ReplaceOne时，更新为当前时间。目前仅支持int64和time.Time类型。

若需要启用软删除特性，那么可以：
``` go
type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
	DeleteAt *int64             `mongox:"delete_at" bson:"delete_at,omitempty"`
}

```
若包含mongox:"delete_at"这个tag，则表示启用软删除。软删除字段目前仅支持* int64和* time.Time类型

当前，您可以使用我们提供的fields.Field字段：
``` go
type User struct {
	fields.Field `bson:"inline"`

	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
}

```

注意：
+ 可以包含多个update_at和delete_at字段
+ 支持解析匿名结构体里的update_at和delete_at，但必须得有bson:"inline"这个tag

## 注册model
您暂时必须先注册model，才能正常使用update_at和delete_at特性
``` go
m.RegisterModels(
	new(User),
	new(Role),
)

```

## InsertOne
您可以传入model，而无需指定collection
``` go
func (m *Mongo) InsertOne(ctx context.Context, bean any, options ...optionx.InsertOneOption) (result *mongo.InsertOneResult, err error)

```

## InsertMany
您可以传入[]model，而无需指定collection
``` go
func (m *Mongo) InsertMany(ctx context.Context, beans []any, options ...optionx.InsertManyOption) (result *mongo.InsertManyResult, err error)

```

## UpdateOne
``` go
func (m *Mongo) UpdateOne(ctx context.Context, collection string, update any, options ...optionx.UpdateOneOption) (result *mongo.UpdateResult, err error)

```
若需要携带filter，那么可以：
``` go
m.UpdateOne(context.Background(), "my_collection", bson.M{"$set": bson.M{"name": "lilith"}}, optionx.WithFilter(bson.M{"name": "lls"})) // 更新 name = lls 的文档
m.UpdateOne(context.Background(), "my_collection", bson.M{"$set": bson.M{"name": "lilith"}}, optionx.WithId(5)) // 更新 id = 5 的文档

m.UpdateOne(context.Background(), "my_collection", bson.M{"$set": bson.M{"name": "lilith"}}, optionx.WithId(5), optionx.WithUnscoped()) // 更新 id = 5 的文档，忽略软删除特性
```
注意：
+ 为保证update_at和delete_at的字段能正常生效，参数update以及filter请使用bson.M, *bson.M, bson.D, *bson.D或它们的别名中的其中一个

## UpdateMany
施工中...

