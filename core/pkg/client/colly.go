package client

import (
	"core/pkg/log"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"go.uber.org/zap"
	"net"
	"net/http"
	"sync"
	"time"
)

var collector *colly.Collector
var collyOnce = sync.Once{}

func GetColly() *colly.Collector {
	collyOnce.Do(func() {
		collector = newCollector()
	})
	return collector.Clone()
}

// New collector
func newCollector() *colly.Collector {
	c := colly.NewCollector()
	c.SetRequestTimeout(50 * time.Second)
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true, // Colly uses HTTP keep-alive to enhance scraping speed
		DialContext: (&net.Dialer{
			Timeout:   90 * time.Second,
			KeepAlive: 90 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   90 * time.Second,
		ExpectContinueTimeout: 90 * time.Second,
	})

	// 对于匹配的域名(当前配置为任何域名),将请求并发数配置为2
	// 通过测试发现,RandomDelay参数对于同步模式也生效
	if err := c.Limit(&colly.LimitRule{
		// glob模式匹配域名
		// DomainGlob: ,

		// 匹配到的域名的并发请求数
		Parallelism: 5,
		// 在发起一个新请求时的随机等待时间
		RandomDelay: time.Duration(500) * time.Millisecond,
	}); err != nil {
		log.Error("生成一个collector对象, 限速配置失败", zap.Error(err))
	}

	// 是否允许重复请求相同url
	c.AllowURLRevisit = false
	c.Async = false
	c.DetectCharset = true

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		// r.Request.Retry()
		err = r.Request.Retry()
		log.Error("Error occurs", zap.Error(err), zap.Int("statusCode", r.StatusCode), zap.String("url", r.Request.URL.String()))
	})

	c.OnRequest(func(r *colly.Request) {
		log.Info("[Visiting] " + r.URL.String())
	})

	// 随机设置
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	return c
}
