package client

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"testing"
)

func TestColly(t *testing.T) {
	// Instantiate default collector
	c := GetColly()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(response *colly.Response, err error) {
		println(err)
	})

	// Start scraping on https://hackerspaces.org
	err := c.Visit("http://m.xinbanzhu.net/sort/3_1/")
	if err != nil {
		panic(err)
	}
	print("last.........")
}
