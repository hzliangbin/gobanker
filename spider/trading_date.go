package spider

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"gobanker/util"
	"net/http"
)

type TradingDate struct {
	Zrxh int `json:"zrxh"`
	Jybz string `json:"jybz"`
	Jyrq string `json:"jyrq"`
}
type TradingDateMsg struct {
	Data []TradingDate `json:"data"`
	Nowdate string `json:"nowdate"`
}

const baseUrl = "http://www.szse.cn/api/report/exchange/onepersistenthour/monthList"
//TODO 每月1号更新一次，存入数据库后，其它时候直接从数据库取
func TradingDateSpider(month string) error {

	c := colly.NewCollector(colly.Async(true))
	extensions.RandomUserAgent(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Referer","http://www.szse.cn/disclosure/index.html")
		fmt.Printf("Visiting %s", request.URL)
		fmt.Println("")
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

	msg := TradingDateMsg{}
	c.OnResponse(func(response *colly.Response) {
		err := json.Unmarshal(response.Body, &msg)
		if err != nil {
			fmt.Println("反序列化失败")
			fmt.Println(err)
		}
	})

	url := fmt.Sprintf("%s?month=%s", baseUrl, month)
	err := c.Visit(url)
	if err != nil {
		fmt.Println("visit error")
	}

	c.Wait()


	ctx, cli := util.ConnectWithColl("t_trading_date")
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	if err = cli.DropCollection(ctx); err != nil {
		fmt.Println("delete tradingdate collection error")
	}
	if _, err = cli.Collection.InsertMany(ctx, msg.Data); err != nil {
		fmt.Println("insert tradingdate collection error")
		fmt.Println(err)
	}

	return nil
}
