package job

import (
	"core/pkg/common"
	"core/pkg/job/parser"
)

var parserMap = make(map[string]parser.Parser)
var jobMap = make(map[string]Job)

func RegisterParser(parserName string, parser parser.Parser) {
	parserMap[parserName] = parser
}

func GetParser(parserName string) parser.Parser {
	return parserMap[parserName]
}

func RegisterJob(jobName string, job Job) {
	jobMap[jobName] = job
}

func GetJob(jobName string) Job {
	return jobMap[jobName]
}

func init() {
	RegisterParser(common.HomeParser, &parser.HomeParser{})
	RegisterParser(common.CatalogHomeParser, &parser.CatalogHomeParser{})
	RegisterParser(common.CatalogPageParser, &parser.CatalogPageParser{})

	RegisterJob(common.HomeJob, &HomeJob{})
	RegisterJob(common.CatalogHomeJob, &CatalogHomeJob{})
	RegisterJob(common.CatalogPageJob, &CatalogPageJob{})
}
