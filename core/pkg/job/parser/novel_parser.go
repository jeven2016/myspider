package parser

import (
	"core/pkg/client"
	"core/pkg/log"
	"core/pkg/service"
	"core/pkg/stream/message"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

type NovelParser struct {
}

func (np *NovelParser) Parse(params *ParseParams) *ParseResult {
	collector := client.CloneColly()
	if collector == nil {
		return &ParseResult{Payload: nil}
	}

	siteRule := service.GetSiteRule(params.SiteKey)
	if siteRule == nil || siteRule.Rule == nil || siteRule.Rule.Novel == nil {
		log.SugaredLogger().Error("no rule found for site key %v", params.SiteKey)
		return &ParseResult{Payload: nil}
	}

	site := service.GetSiteByKey(params.SiteKey)
	baseUrl, err := site.BaseUrl()
	if err != nil {
		log.SugaredLogger().Error("failed to find the base of home url %v (%v)", params.SiteKey, site.HomeUrl)
		return &ParseResult{Payload: nil}
	}

	//解析书本信息并保存

	var chapterMsg []*message.ChapterMessage

	collector.OnHTML(siteRule.Rule.Home.CatalogPageSelector, func(element *colly.HTMLElement) {
		chapterMsg = np.analyse(element, baseUrl, chapterMsg, params)
	})

	err = collector.Visit(params.Url)
	if err != nil {
		log.SugaredLogger().Error("failed to visit novel page", zap.String("siteKey", params.SiteKey),
			zap.String("url", params.Url))
	}
	return &ParseResult{Payload: chapterMsg}
}

func (ch *NovelParser) analyse(element *colly.HTMLElement, baseUrl string,
	chapterMsg []*message.ChapterMessage,
	params *ParseParams) []*message.ChapterMessage {
	href := element.Attr("href")

	return chapterMsg
}
