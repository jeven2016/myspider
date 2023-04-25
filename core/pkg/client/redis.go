package client

import (
	"context"
	"core/pkg/config"
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

func GetRedisClient() *redis.Client {
	return redisClient.Client
}
