package parser

import (
	"core/pkg/client"
	"core/pkg/common/utils"
	"core/pkg/log"
	"core/pkg/service"
	"core/pkg/stream/message"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type CatalogHomeParser struct {
}

func (ch *CatalogHomeParser) Parse(params *ParseParams) *ParseResult {
	collector := client.CloneColly()
	if collector == nil {
		return &ParseResult{Payload: nil}
	}

	siteRule := service.GetSiteRule(params.SiteKey)
	if siteRule == nil || siteRule.Rule == nil || siteRule.Rule.Home == nil {
		log.SugaredLogger().Error("no rule found for site key %v", params.SiteKey)
		return &ParseResult{Payload: nil}
	}

	site := service.GetSiteByKey(params.SiteKey)
	baseUrl, err := site.BaseUrl()
	if err != nil {
		log.SugaredLogger().Error("failed to find the base of home url %v (%v)", params.SiteKey, site.HomeUrl)
		return &ParseResult{Payload: nil}
	}

	var scPageMsgs []*message.SiteCatalogPageMessage

	collector.OnHTML(siteRule.Rule.Home.CatalogPageSelector, func(element *colly.HTMLElement) {
		scPageMsgs = ch.analyse(element, baseUrl, scPageMsgs, params)
	})

	err = collector.Visit(params.Url)
	if err != nil {
		log.SugaredLogger().Error("failed to visit catalog page", zap.String("siteKey", params.SiteKey),
			zap.String("url", params.Url))
	}
	return &ParseResult{Payload: scPageMsgs}
}

func (ch *CatalogHomeParser) analyse(element *colly.HTMLElement, baseUrl string, scPageMsgs []*message.SiteCatalogPageMessage, params *ParseParams) []*message.SiteCatalogPageMessage {
	href := element.Attr("href")
	hasHttpPrefix := utils.CheckHttpUrl(href)
	if len(href) != 0 {
		compile, err := regexp.Compile(".*/\\d+_(\\d+)/?$")
		if err != nil {
			log.Logger().Error("failed to analyse the url", zap.String("href", href))
			return nil
		}
		subStrings := compile.FindStringSubmatch(href)
		if len(subStrings) > 1 {
			lastPageNo, err := convertor.ToInt(subStrings[1])
			if err != nil {
				log.Logger().Error("failed to convert string to int", zap.String("value", subStrings[1]), zap.Error(err))
				return nil
			}

			lpnInt := int(lastPageNo)
			compile, err = regexp.Compile("(.*/\\d+)_\\d+/?$")

			//the page index starts from 1 instead of zero
			for i := 1; i <= lpnInt; i++ {
				urlPrefixSlice := compile.FindStringSubmatch(href)
				if len(urlPrefixSlice) > 1 {
					nextPageLink := fmt.Sprintf("%v_%v/", urlPrefixSlice[1], i)

					if !hasHttpPrefix {
						//append the base url
						nextPageLink = baseUrl + "/" + strings.TrimLeft(nextPageLink, "/")
					}

					scPageMsgs = append(scPageMsgs, &message.SiteCatalogPageMessage{
						SiteKey:         params.SiteKey,
						Page:            i,
						CatalogPageLink: nextPageLink,
					})
				}

			}
		}
	}
	return scPageMsgs
}
