package models

import (
	"github.com/astaxie/beego/orm"
)

type Bank struct {
	Code string `orm:"pk"`
	Name string
	Category int //分类，0国有，1股份，2城商，3农村
//	EstablishAt time.Time
//	headquarters string
}

func init() {
	orm.RegisterModel(new(Bank))
}
