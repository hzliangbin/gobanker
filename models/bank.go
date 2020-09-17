package models

import "github.com/astaxie/beego/orm"

type Bank struct {
	Code string `orm:"pk"`
	Name string
}

func init() {
	orm.RegisterModel(new(Bank))
}
