package job

import (
	"context"
	"core/pkg/common"
	"core/pkg/job/parser"
	"core/pkg/model"
	"errors"
)

type ExecutionResult struct {
	JobName string
	Err     error
}

type Job interface {
	Validate(ctx context.Context) error
	ConvertParams(jsonData string) *parser.ParseParams
	Run(ctx context.Context)
}

type WrapperJob struct {
	JobParams  *model.SiteJob
	JobHandler Job
	Parser     parser.Parser
}

func NewWrapperJob(jobParams *model.SiteJob) (*WrapperJob, error) {
	realJob := GetJob(jobParams.Type)
	if realJob == nil {
		return nil, errors.New("no job registered for type " + jobParams.Type)
	}

	pageParser := GetParser(jobParams.Parser)
	if pageParser == nil {
		return nil, errors.New("no parser registered for " + jobParams.Parser)
	}
	return &WrapperJob{
		JobParams:  jobParams,
		JobHandler: realJob,
		Parser:     pageParser,
	}, nil
}

func (job *WrapperJob) Launch(ctx context.Context, site *model.Site) {
	newCtx := context.WithValue(ctx, common.ParserCtx, job.Parser)
	newCtx = context.WithValue(newCtx, common.SiteCtx, site)

	errChan := ctx.Value(common.ErrChan).(chan ExecutionResult)

	if err := job.JobHandler.Validate(newCtx); err != nil {
		errChan <- ExecutionResult{JobName: job.JobParams.Name, Err: err}
	}
	go job.JobHandler.Run(newCtx)
}
