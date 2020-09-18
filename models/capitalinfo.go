package models

import "github.com/astaxie/beego/orm"

type CapitalInfo struct {
	Code string
	Name string
	CoreTier1CAR float64
	Tier1CAR float64
	CAR float64
}

func init()  {
	orm.RegisterModel(new(CapitalInfo))
}
