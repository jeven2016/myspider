package job

import "core/pkg/config"

func PrepareJobs(cfg *config.Config) error {

	if cfg.Execution != nil {
		for key, webSite := range cfg.Execution {
			if !webSite.Enabled {
				continue
			}
			launchJob(key, &webSite)
		}
	}
	return nil

}

func launchJob(key string, webSite *config.WebSite) error {

	for _, job := range webSite.Jobs {
		jobInst = &GenericJob{jobParams: &job}
		//todo
	}

	return nil
}
