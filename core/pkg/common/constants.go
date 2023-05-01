package common

const StreamValuesKey = "data"

// ConfigFileType 配置文件类型
type ConfigFileType int

var JobParallelism uint = 1

const (
	SysCfg ConfigFileType = iota
	SiteRuleCfg
)

const (
	SysCfgName   = "config" //system config
	SiteRuleName = "siterule"
	CfgFormat    = "yaml"
)

var CfgSearchPaths = []string{"/etc/myspider", "./conf", "./"}

const (
	HomeParser        = "HomeParser"
	CatalogHomeParser = "CatalogHomeParser"
	CatalogPageParser = "CatalogPageParser"
	BookParser        = "BookParser"
	ChapterParser     = "ChapterParser"

	HomeJob        = "HomeJob"
	CatalogHomeJob = "CatalogHomeJob"
	CatalogPageJob = "CatalogPageJob"
)

const (
	SiteCtx   = "site"
	ParserCtx = "parser"
)

const ErrChan = "errChan"
