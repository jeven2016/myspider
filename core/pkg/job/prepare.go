package job

import (
	"context"
	"core/pkg/config"
	"core/pkg/log"
)

func PrepareJobs(cfg *config.SysConfig) map[string]error {
	var err error
	var errMap = make(map[string]error)

	if cfg.Execution != nil {
		for key, webSite := range cfg.Execution {
			if !webSite.Enabled {
				continue
			}
			if err = launchJob(key, &webSite); err != nil {
				errMap[key] = err
			}
		}
	}
	return errMap

}

func launchJob(key string, webSite *config.WebSite) error {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer func() {
		if err := recover(); err != nil {
			log.Warn("a job interrupted")
			cancelFunc()
		}
	}()

	for _, j := range webSite.Jobs {
		job := NewGenericJob(&j)
		if j.Enabled {
			if err := job.Launch(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
