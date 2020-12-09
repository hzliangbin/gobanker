package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"regexp"
	"strings"
)

func SinaIndexSpider(queryCode *string) error{
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
	//var msg *string
	c.OnResponse(func(response *colly.Response) {
		//err := json.Unmarshal(response.Body,)
		//fmt.Println(string(response.Body))
		msg := string(response.Body)
		parseDetail(&msg)

	})
	url := fmt.Sprintf("http://hq.sinajs.cn/list=%s", *queryCode)
	err := c.Visit(url)
	if err != nil {
		fmt.Println("visit error")
	}
	c.Wait()

	return nil
}

func parseDetail(content *string) ([]string, error) {
	*content = strings.Trim(*content,"\n")
	items := strings.Split(*content, ";")
	for _, item := range items {
		item = strings.TrimSpace(item)
		item = regexp.MustCompile("^var\\D+|\"").ReplaceAllString(item,"")
		item = strings.ReplaceAll(item,"=",",")
		data := strings.Split(item,",")

		if len(data) < 30 {
			fmt.Println("代码不存在")
		} else {
			fmt.Printf("%v",data)
			fmt.Println("")
		}
	}
	return  nil, nil


}