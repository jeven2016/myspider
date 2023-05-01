package client

import (
	"core/pkg/config"
	"core/pkg/log"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/gocolly/colly/v2/proxy"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

var collector *colly.Collector
var collyOnce sync.Once

var retryLock sync.Locker
var retryUrlsMap = make(map[string]uint8)

func GetColly() *colly.Collector {
	collyOnce.Do(func() {
		newCollector()
	})
	return collector
}

// New collector
func newCollector() {
	c := colly.NewCollector()
	c.SetRequestTimeout(50 * time.Second)
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true, // Colly uses HTTP keep-alive to enhance scraping speed
		DialContext: (&net.Dialer{
			Timeout:   20 * time.Second,
			KeepAlive: 90 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 30 * time.Second,
	})
	spider := config.GetSysConfig().Spider
	if spider != nil && strings.TrimSpace(spider.HttpProxy) != "" {
		proxyList := []string{spider.HttpProxy}

		// Create Rotating Proxy Switcher
		rp, err := proxy.RoundRobinProxySwitcher(proxyList...)
		if err != nil {
			log.SugaredLogger().Error(err)
			return
		}

		// Set Collector To Use Proxy Switcher Function
		c.SetProxyFunc(rp)
	}

	// 对于匹配的域名(当前配置为任何域名),将请求并发数配置为2
	// 通过测试发现,RandomDelay参数对于同步模式也生效
	//if err := c.Limit(&colly.LimitRule{
	//	// glob模式匹配域名
	//	// DomainGlob: ,
	//
	//	// 匹配到的域名的并发请求数
	//	Parallelism: 5,
	//	// 在发起一个新请求时的随机等待时间
	//	RandomDelay: time.Duration(500) * time.Millisecond,
	//}); err != nil {
	//	log.Error("生成一个collector对象, 限速配置失败", zap.Error(err))
	//}

	// 是否允许重复请求相同url
	c.AllowURLRevisit = true
	c.Async = false
	c.DetectCharset = true

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		url := r.Request.URL.String()
		retryLock.Lock()
		defer retryLock.Unlock()

		if _, ok := retryUrlsMap[url]; !ok {
			retryUrlsMap[url]++
			err = r.Request.Retry()
			return
		}

		if retryUrlsMap[url] < 3 {
			retryUrlsMap[url]++
			err = r.Request.Retry()
			return
		}

		if err != nil {
			log.Logger().Error("Error occurs", zap.Error(err), zap.Int("statusCode", r.StatusCode), zap.String("url", r.Request.URL.String()))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Logger().Info("[Visiting] " + r.URL.String())
	})

	// 随机设置
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	collector = c
}

func CloneColly() *colly.Collector {
	clt := GetColly().Clone()
	if clt == nil {
		log.SugaredLogger().Error("no colly connector retrieved")
		return nil
	}
	return clt
}
