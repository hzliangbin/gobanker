package models

type Bank struct {
	Code     string
	Name     string
	Category int //分类，0国有，1股份，2城商，3农村
	//	EstablishAt time.Time
	//	headquarters string
}
