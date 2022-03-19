package models

type ManageInfo struct {
	Id              float64 `orm:"pk"`
	Code            string
	FinQuarter      string
	Name            string
	ROE             float64
	ROA             float64
	PB              float64
	TTM             float64
	NetProfit       float64
	NetProfitGrowth float64
	Revenue         float64
	RevenueGrowth   float64
}
