package models

import "github.com/astaxie/beego/orm"

type AssetsInfo struct {
	Code string
	Name string
	TotalAssets float64
	NetAssets float64
	TotalAssetsGrowth float64
	NetAssetsGrowth float64
}

func init()  {
	orm.RegisterModel(new(AssetsInfo))
}

