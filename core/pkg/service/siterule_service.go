package service

import (
	"core/pkg/config"
	"core/pkg/model"
)

func GetSiteRule(siteKey string) *model.SiteRule {
	var ruleConfig *model.SiteRule
	if ruleConfig = config.GetSiteRuleConfig(); ruleConfig == nil {
		return ruleConfig
	}

	if ruleConfig.Key == siteKey {
		return ruleConfig
	}
	return nil
}
