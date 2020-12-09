package main

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "gobanker/routers"
	"gobanker/spider"
)

func init() {
	//orm.RegisterDriver("postgres",orm.DRPostgres)
	//var dataSource string = beego.AppConfig.String("pgsql_datasource")
	//orm.RegisterDataBase("default","postgres", dataSource)
	//orm.RunSyncdb("default",false,true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//spider.CsIndexIndustryHandler()

	//spider.TradingDateSpider("2020-12")
	codes := []string{"sz000002","sh600519"}
	var buffer bytes.Buffer
	for _, code := range codes {
		buffer.WriteString(code)
		buffer.WriteString(",")
	}
	queryCodes := buffer.String()

	fmt.Println(queryCodes)
	if err := spider.SinaIndexSpider(&queryCodes); err != nil {
		fmt.Println("sina index ok")
	}
	orm.Debug = true
	beego.Run()
}
