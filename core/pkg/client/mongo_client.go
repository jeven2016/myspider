package client

import (
	"context"
	"core/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var mongoClient *MongoClient
var once = sync.Once{}

type MongoClient struct {
	Client *mongo.Client
	Db     *mongo.Database
	Config *config.MongoDb
}

func GetMongoClient(cfg *config.MongoDb) *MongoClient {
	return mongoClient
}

func InitMongodbClient(mongodbCfg *config.MongoDb) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbCfg.Uri))
	if err != nil {
		return err
	}

	// 检测MongoDB是否连接成功
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	//初始化全局Db
	db := client.Database(mongodbCfg.Database)

	mongoClient = &MongoClient{Db: db, Client: client}
	return err
}
