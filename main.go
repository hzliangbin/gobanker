package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gobanker/models"
	_ "gobanker/routers"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("postgres",orm.DRPostgres)
	orm.RegisterDataBase("default","postgres","user=postgres password=Gobanker@2020 dbname=gobanker host=114.67.170.73 sslmode=disable")
	orm.RunSyncdb("default",false,true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	bank := new(models.Bank)
	bank.Code = "600036"
	bank.Name = "招商银行"
	fmt.Println(o.Insert(bank))
	beego.Run()
}
