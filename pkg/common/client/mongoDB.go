package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	cfg "qx/pkg/common/config"
)

// InitMongoDB ...
func InitMongoDB(config cfg.MongoDBConfiguration) (*mongo.Database, error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(config.Host)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	mongoDB := client.Database(config.DBName)
	return mongoDB, nil
}
