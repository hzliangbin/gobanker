package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "gobanker/routers"
	"gobanker/spider"
	"fmt"
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
	fmt.Println("teston heiheihei")
	spider.TradingDateSpider("2020-12")
	orm.Debug = true
	beego.Run()
}
