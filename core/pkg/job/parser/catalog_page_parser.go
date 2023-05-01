package parser

import (
	"core/pkg/client"
	"core/pkg/common/utils"
	"core/pkg/log"
	"core/pkg/model"
	"core/pkg/service"
	"core/pkg/stream/message"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strings"
)

type CatalogPageParser struct {
}

func (cp *CatalogPageParser) Parse(params *ParseParams) *ParseResult {
	collector := client.CloneColly()
	if collector == nil {
		return &ParseResult{Payload: nil}
	}

	//parse the books listed on current page
	siteRule := service.GetSiteRule(params.SiteKey)
	if siteRule == nil || siteRule.Rule == nil || siteRule.Rule.Book == nil {
		log.SugaredLogger().Error("no rule found for parsing book: site %v", params.SiteKey)
		return &ParseResult{Payload: nil}
	}

	site := service.GetSiteByKey(params.SiteKey)
	baseUrl, err := site.BaseUrl()
	if err != nil {
		log.SugaredLogger().Error("failed to find the base of home url %v (%v)", params.SiteKey, site.HomeUrl)
		return &ParseResult{Payload: nil}
	}

	var bookMessages []*message.BookMessage
	collector.OnHTML(siteRule.Rule.Book.BookUrlSelector, func(element *colly.HTMLElement) {
		bookMessages = cp.analyse(element, baseUrl, bookMessages, params, siteRule.Rule.Book)
	})

	err = collector.Visit(params.Url)
	if err != nil {
		log.Logger().Error("failed to visit individual catalog page",
			zap.String("siteKey", params.SiteKey),
			zap.String("url", params.Url))
	}
	return &ParseResult{Payload: bookMessages}
}

func (cp *CatalogPageParser) analyse(element *colly.HTMLElement, baseUrl string, messages []*message.BookMessage, params *ParseParams, bookRule *model.BookRule) []*message.BookMessage {
	name := element.DOM.Text()

	link, exists := element.DOM.Attr(bookRule.BookUrlAttr)
	if !exists || len(link) == 0 {
		log.SugaredLogger().Warn("book [%v] doesn't has link %v", name, link)
		return nil
	}
	hasHttpPrefix := utils.CheckHttpUrl(link)
	if !hasHttpPrefix {
		//append the base url
		link = baseUrl + "/" + strings.TrimLeft(link, "/")
	}

	messages = append(messages, &message.BookMessage{
		SiteKey:   params.SiteKey,
		CatalogId: "TODO",
		Name:      name,
		BookLink:  link,
	})

	return messages
}
