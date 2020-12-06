package spider

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
)

func SinaIndexSpider(stockCode *[]string) error{
	c := colly.NewCollector(colly.Async(true))
	extensions.RandomUserAgent(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnRequest(func(request *colly.Request) {
		//request.Headers.Set("Referer",)
	})
	return nil
}