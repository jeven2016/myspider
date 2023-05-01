package parser

import (
	"core/pkg/client"
	"core/pkg/common/utils"
	"core/pkg/log"
	"core/pkg/service"
	"core/pkg/stream/message"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strings"
)

type HomeParser struct {
}

func (hp *HomeParser) Parse(params *ParseParams) *ParseResult {
	collector := client.CloneColly()
	if collector == nil {
		return &ParseResult{Payload: nil}
	}

	siteRule := service.GetSiteRule(params.SiteKey)
	if siteRule == nil || siteRule.Rule == nil || siteRule.Rule.Home == nil {
		log.SugaredLogger().Error("no rule found for site key %v", params.SiteKey)
		return &ParseResult{Payload: nil}
	}

	homeRule := siteRule.Rule.Home

	var catalogMsgs []*message.SiteCatalogHomeMessage

	collector.OnHTML(homeRule.CatalogTitleSelector, func(element *colly.HTMLElement) {
		name := element.DOM.Text()
		link, exists := element.DOM.Attr(homeRule.CatalogUrlAttr)
		link = strings.Trim(link, " ")
		if !exists || len(link) == 0 {
			log.SugaredLogger().Warn("[%v] doesn't has link %v", name, link)
			return
		}

		hasHttpPrefix := utils.CheckHttpUrl(params.Url)
		if !hasHttpPrefix {
			// get prefix from site home url
			site := service.GetSiteByKey(params.SiteKey)
			siteBaseUrl, err := site.BaseUrl()
			if err != nil {
				log.SugaredLogger().Error("failed to find the base of CatalogHome url %v (%v)", params.SiteKey, site.HomeUrl)
				return
			}
			link = siteBaseUrl + "/" + strings.TrimLeft(link, "/")
		} else {
			// get prefix from the link itself
			baseUrl, err := utils.GetBaseUrl(params.Url)
			if err != nil {
				log.Logger().Error("invalid url format", zap.String("homeUrl", params.Url), zap.Error(err))
				return
			}
			link = baseUrl + "/" + strings.TrimLeft(link, "/")
		}

		catalogMessage := &message.SiteCatalogHomeMessage{SiteKey: params.SiteKey, Name: name, CatalogLink: link}
		catalogMsgs = append(catalogMsgs, catalogMessage)
	})

	err := collector.Visit(params.Url)
	if err != nil {
		log.SugaredLogger().Error("failed to visit", zap.String("url", params.Url))
	}
	return &ParseResult{Payload: catalogMsgs}
}
