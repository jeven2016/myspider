package job

import (
	"context"
	"core/pkg/client"
	"core/pkg/common"
	"core/pkg/job/parser"
	"core/pkg/stream"
	"github.com/reugn/go-streams/flow"
)

func ParseLinks(pageParser parser.Parser, j *parser.ParseParams) interface{} {
	if j == nil {
		return nil
	}
	result := pageParser.Parse(j)
	return result.Payload
}

func CreateSource(ctx context.Context, streamName string, consumerGroupName string) (*client.RedisClient,
	*stream.RedisStreamSource, error) {
	errChan := ctx.Value(common.ErrChan).(chan ExecutionResult)

	//construct source
	redisClient := client.GetRedisClient()
	source, err := stream.NewRedisStreamSource(ctx, redisClient,
		streamName, consumerGroupName)
	if err != nil {
		errChan <- ExecutionResult{JobName: common.HomeJob, Err: err}
	}
	return redisClient, source, err
}

func FlatMany[T any](pageParser parser.Parser) *flow.FlatMap[*parser.ParseParams, T] {
	return flow.NewFlatMap(func(j *parser.ParseParams) []T {
		return ParseLinks(pageParser, j).([]T)
	}, common.JobParallelism)
}
