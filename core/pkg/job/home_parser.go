package job

import (
	"core/pkg/client"
	"core/pkg/common"
	"github.com/gocolly/colly/v2"
)

type HomeParser struct {
}

func (hp *HomeParser) Parse(params *ParseParams) ([]string, error) {
	collector := client.GetColly()
	ctx := colly.NewContext()

	collector.OnHTML("body", func(element *colly.HTMLElement) {

	})

	err := collector.Request("GET", params.Url, nil, ctx, nil)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}

func init() {
	RegisterParser(common.HomeParser, &HomeParser{})
}
