package job

import (
	"context"
	"core/pkg/client"
	"core/pkg/common"
	"core/pkg/config"
	"core/pkg/stream"
	"errors"
	"github.com/reugn/go-streams/flow"
)

type Job interface {
	preExecute()
	Run()
}

type GenericJob struct {
	JobParams     *config.JobParams
	enabled       bool
	enableLoad    bool
	enableSave    bool
	enableAnalyse bool
}

func NewGenericJob(jobParams *config.JobParams) *GenericJob {
	var load bool
	var save bool
	var analyse bool

	if jobParams.Action == nil {
		load = true
		save = true
		analyse = true
	} else {
		action := jobParams.Action
		load = action.Load
		save = action.Save
		analyse = action.Analyse
	}

	return &GenericJob{
		JobParams:     jobParams,
		enabled:       true,
		enableLoad:    load,
		enableSave:    save,
		enableAnalyse: analyse,
	}
}

func (job *GenericJob) Load(context.Context) (bool, error) {
	if !job.enableLoad || !job.enabled {
		return false, nil
	}
	return true, nil
}

func (job *GenericJob) Analyse() {

}

func (job *GenericJob) Launch(ctx context.Context) error {
	redisClient := client.GetRedisClient()
	source, err := stream.NewRedisStreamSource(ctx, redisClient,
		common.HomeUrlStream, common.HomeUrlStreamConsumer)
	if err != nil {
		return err
	}

	if len(job.JobParams.Url) == 0 {
		return errors.New("url is required for paring site's home page")
	}
	parser := GetParser(job.JobParams.Parser)
	if parser == nil {
		return errors.New("invalid parser for paring site's home page")
	}

	converter := func(homeUrl string) *config.JobParams {
		return nil
	}

	flow.NewMap(converter, 1)
	parseFlow := flow.NewFlatMap(parser.Parse, 1)

	sink := stream.NewRedisStreamSink(ctx, redisClient, common.SiteCatalogUrlStream)

	source.
		Via(parseFlow).
		To(sink)
}
