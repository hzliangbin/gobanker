package models

import "github.com/astaxie/beego/orm"

type CapitalInfo struct {
	Id float64 `orm:"pk"`
	Code string
	Name string
	FinQuarter string
	CoreTier1CAR float64
	Tier1CAR float64
	CAR float64
}

func init()  {
	orm.RegisterModel(new(CapitalInfo))
}
