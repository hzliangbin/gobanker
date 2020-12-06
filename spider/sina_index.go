package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
)

func SinaIndexSpider(stockCode *[]string) error{
	c := colly.NewCollector(colly.Async(true))
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnRequest(func(request *colly.Request) {
		//request.Headers.Set("Referer",)
	})

	retryCount := 0
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Sth went wrong:",err)
		if retryCount < 3 {
			retryCount += 1
			err := response.Request.Retry()
			fmt.Println("retry wrong", err)
		}
	})

	c.OnResponse(func(response *colly.Response) {
		//err := json.Unmarshal(response.Body,)
		fmt.Println(string(response.Body))
	})
	url := fmt.Sprintf("http://hq.sinajs.cn/list=%s", *stockCode)
	err := c.Visit(url)
	if err != nil {
		fmt.Println("visit error")
	}
	c.Wait()
	return nil
}