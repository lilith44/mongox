# mongox
基于官方包go.mongodb.org/mongo-driver进行简单封装，主要有以下特性：

+ 支持软删除
+ 支持动态迁移collection以及其index
+ 支持了日志
+ 使用选项设计模式，减少了部分方法的参数

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
