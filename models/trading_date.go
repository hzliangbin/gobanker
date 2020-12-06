package models

type TradingDate struct {
	Zrxh int `json:"zrxh"`
	Jybz string `json:"jybz"`
	Jyrq string `json:"jyrq"`
}
type TradingDateMsg struct {
	Data []TradingDate `json:"data"`
	Nowdate string `json:"nowdate"`
}

