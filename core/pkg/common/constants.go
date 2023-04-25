package common

// ConfigFileType 配置文件类型
type ConfigFileType int

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
	CatalogParser     = "CatalogParser"
	CatalogPageParser = "CatalogPageParser"
	BookParser        = "BookParser"
	ChapterParser     = "ChapterParser"
)

const (
	HomeUrlStream         = "homeUrlStream"
	HomeUrlStreamConsumer = "HomeUrlStreamConsumer"

	SiteCatalogUrlStream         = "SiteCatalogUrlStream"
	SiteCatalogUrlStreamConsumer = "SiteCatalogUrlStreamConsumer"
)
