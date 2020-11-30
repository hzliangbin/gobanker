package main

import (
	"github.com/astaxie/beego/orm"
	_ "gobanker/routers"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
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
	orm.Debug = true
	beego.Run()
}
