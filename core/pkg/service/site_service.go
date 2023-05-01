package service

import (
	"core/pkg/config"
	"core/pkg/model"
)

func GetSiteByKey(siteKey string) *model.Site {
	if siteKey == "" {
		return nil
	}
	// get from config
	site, ok := config.GetSysConfig().Execution[siteKey]
	if ok {
		return site
	}

	return nil
}
