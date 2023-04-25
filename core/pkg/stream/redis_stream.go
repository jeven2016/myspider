package stream

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/reugn/go-streams"
	"github.com/reugn/go-streams/flow"
	"log"
	"math/rand"
	"strconv"
)

// this file forked and enhanced based on https://github.com/reugn/go-streams since I want to use go-redis/v9 and redis stream feature

// RedisStreamSource is a Redis Pub/Sub Source
type RedisStreamSource struct {
	ctx           context.Context
	redisClient   *redis.Client
	out           chan interface{}
	streamName    string
	consumerGroup string
}

// NewRedisStreamSource returns a new RedisStreamSource instance
func NewRedisStreamSource(ctx context.Context, client *redis.Client, streamName string,
	consumerGroup string) (*RedisStreamSource, error) {
	var err error
	if err = ensureConsumeGroup(client, consumerGroup, streamName); err != nil {
		return nil, err
	}

	source := &RedisStreamSource{
		ctx:           ctx,
		redisClient:   client,
		out:           make(chan interface{}),
		streamName:    streamName,
		consumerGroup: consumerGroup,
	}

	go source.init(ctx)
	return source, nil
}

func ensureConsumeGroup(client *redis.Client, consumerGroup, stream string) error {
	var groups []redis.XInfoGroup
	var err error

	//当无法获取到group信息时，创建一个消费group
	if groups, err = client.XInfoGroups(context.Background(), stream).Result(); err != nil {
		var groupExists bool
		for _, g := range groups {
			if g.Name == consumerGroup {
				return nil
			}
		}

		if !groupExists {
			//You can use the XGROUP CREATE command with MKSTREAM option, to create an empty stream
			//XGroupCreate 方法要求先有stream的存在才能创建group
			return client.XGroupCreateMkStream(context.Background(), stream, consumerGroup, "0").Err()
		}
	}

	return nil
}

// init starts the main loop
func (rs *RedisStreamSource) init(ctx context.Context) {
	defer func() {
		close(rs.out)
		rs.redisClient.Close()
	}()

loop:
	for {

		select {
		case <-ctx.Done():
			break loop
		default:
			consumer := rs.streamName + ":consumer:" + strconv.Itoa(rand.Int())
			entries, err := rs.redisClient.XReadGroup(context.Background(), &redis.XReadGroupArgs{
				Group:    rs.consumerGroup,
				Consumer: consumer,

				Streams: []string{rs.streamName, ">"},
				Count:   2,
				Block:   0,
			}).Result()
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(entries[0].Messages); i++ {
				messageID := entries[0].Messages[i].ID
				values := entries[0].Messages[i].Values
				rs.out <- values
				rs.redisClient.XAck(context.Background(), rs.streamName, rs.consumerGroup, messageID)

			}
		}

	}

}

// Via streams data through the given flow
func (rs *RedisStreamSource) Via(_flow streams.Flow) streams.Flow {
	flow.DoStream(rs, _flow)
	return _flow
}

// Out returns an output channel for sending data
func (rs *RedisStreamSource) Out() <-chan interface{} {
	return rs.out
}

// RedisStreamSink is a Redis Pub/Sub Sink
type RedisStreamSink struct {
	redisClient *redis.Client
	in          chan interface{}
	streamName  string
}

// NewRedisStreamSink returns a new RedisStreamSink instance
func NewRedisStreamSink(ctx context.Context, client *redis.Client, streamName string) *RedisStreamSink {
	sink := &RedisStreamSink{
		client,
		make(chan interface{}),
		streamName,
	}

	go sink.init(ctx)
	return sink
}

// init starts the main loop
func (rs *RedisStreamSink) init(ctx context.Context) {
	defer rs.redisClient.Close()
	for msg := range rs.in {
		err := rs.redisClient.XAdd(ctx, &redis.XAddArgs{Stream: rs.streamName,
			NoMkStream: false,  // * 默认false,当为false时,key不存在，会新建
			MaxLen:     100000, // * 指定stream的最大长度,当队列长度超过上限后，旧消息会被删除，只保留固定长度的新消息
			Approx:     false,  // * 默认false,当为true时,模糊指定stream的长度
			ID:         "*",    // 消息 id，我们使用 * 表示由 redis 生成
			// MinID: "id",            // * 超过阈值，丢弃设置的小于MinID消息id【基本不用】
			// Limit: 1000,            // * 限制长度【基本不用】
			Values: msg}).Err()

		if err != nil {
			log.Printf("NewRedisStreamSink.init():Send message failed with: %s", err)
		}
	}

}

// In returns an input channel for receiving data
func (rs *RedisStreamSink) In() chan<- interface{} {
	return rs.in
}
