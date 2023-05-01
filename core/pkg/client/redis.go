package client

import (
	"context"
	"core/pkg/common"
	"core/pkg/config"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisClient *RedisClient

type RedisClient struct {
	Client *redis.Client
}

func InitRedisClient(redisCfg *config.Redis) error {
	client := redis.NewClient(&redis.Options{
		Addr:         redisCfg.Address,
		Password:     redisCfg.Password,
		DB:           redisCfg.Db,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  time.Duration(redisCfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(redisCfg.WriteTimeout) * time.Second,
		PoolSize:     redisCfg.PoolSize,
		PoolTimeout:  time.Duration(redisCfg.PoolTimeout) * time.Second,
	})

	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		return err
	}
	redisClient = &RedisClient{client}
	return nil
}

func GetRedisClient() *RedisClient {
	return redisClient
}

func (c *RedisClient) PublishMessage(ctx context.Context, data interface{}, streamName string) error {
	if data == nil {
		return errors.New(fmt.Sprintf("cannot publis empty data, stream is %v", streamName))
	}
	json, err := convertor.ToJson(data)
	if err != nil {
		return common.JsonConvertErr
	}

	//just send the json data into stream since it's too complicated to map a struct to map, there would be
	//different kind of exceptions that need to handle
	err = c.Client.XAdd(ctx, &redis.XAddArgs{Stream: streamName,
		NoMkStream: false,  // * 默认false,当为false时,key不存在，会新建
		MaxLen:     100000, // * 指定stream的最大长度,当队列长度超过上限后，旧消息会被删除，只保留固定长度的新消息
		Approx:     false,  // * 默认false,当为true时,模糊指定stream的长度
		ID:         "*",    // 消息 id，我们使用 * 表示由 redis 生成
		// MinID: "id",            // * 超过阈值，丢弃设置的小于MinID消息id【基本不用】
		// Limit: 1000,            // * 限制长度【基本不用】
		Values: map[string]any{
			common.StreamValuesKey: json,
		}}).Err()

	return err
}
