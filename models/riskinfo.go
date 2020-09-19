package models

type RiskInfo struct {
	Id float64 `orm:"pk"`
	Code string
	Name string
	FinQuarter string
	NPLRatio float64  //不良发生率
	NPLCoverageRatio float64  //不良拨备覆盖率
}
