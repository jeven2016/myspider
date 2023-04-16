package job

import "core/pkg/config"

type Job interface {
	Load()
	Analyse()
}

type GenericJob struct {
	jobParams *config.JobParams
}

func (job *GenericJob) Load() {

}

func (job *GenericJob) Analyse() {

}
