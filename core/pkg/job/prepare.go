package job

import (
	"context"
	"core/pkg/common"
	"core/pkg/config"
	"core/pkg/log"
	"core/pkg/model"
	"go.uber.org/zap"
)

func PrepareJobs(cfg *config.SysConfig) {
	var err error
	var errMap = make(map[string]error)

	if cfg.Execution != nil {
		for key, site := range cfg.Execution {
			if !site.Enabled {
				continue
			}
			if err = launchJob(site); err != nil {
				errMap[key] = err
			}
		}
	}
	if len(errMap) > 0 {
		for k, v := range errMap {
			log.SugaredLogger().Errorf("%s, %v", k, v)
		}
	}

}

func launchJob(site *model.Site) error {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer func() {
		if err := recover(); err != nil {
			log.SugaredLogger().Warn("a job to launch is abnormally interrupted  ", zap.Any("error", err))
			cancelFunc()
		}
	}()
	errChan := make(chan ExecutionResult, 100)
	ctx = context.WithValue(ctx, common.ErrChan, errChan)

	for _, j := range site.Jobs {
		job, err := NewWrapperJob(&j)
		if err != nil {
			return err
		}
		job.Launch(ctx, site)
	}

	// another goroutine to handle errors occurred in job
	go func() {
		var result ExecutionResult

		// print the errors of job
		for {
			select {
			case result = <-errChan:
				log.Logger().Error("job failed", zap.String("jobName", result.JobName),
					zap.Error(result.Err))

			case <-ctx.Done():
				// job cancelled
				log.SugaredLogger().Info("Jobs cancelled")
				return
			}
		}

	}()

	return nil
}
