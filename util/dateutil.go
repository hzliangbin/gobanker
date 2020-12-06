package util

import (
	"github.com/hzliangbin/date"
	"go.mongodb.org/mongo-driver/bson"
	"gobanker/models"
)



func GetDateYYYYMMdd() (string) {
	d := date.Now()
	res, err := d.Format("yyyyMMdd",false)
	if err != nil {
		return ""
	}
	return res
}

func IsTradingDay() (bool, error) {
	d := date.Now()
	//获取本月月份
	today, err := d.Format("yyyy-MM-dd",false)
	//month, err := d.Format("yyyy-MM",false)
	if err != nil {
		return false, err
	}
	//从数据库中获取trading_date表数据，有则直接返回，没有则更新当月数据再返回
	ctx, cli := ConnectWithColl("t_trading_date")
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	one := models.TradingDate{}
	err = cli.Find(ctx,bson.M{"jyrq":today}).One(&one);
	if err == nil {
		return one.Jybz == "1", nil
	}
	//if err = spider.TradingDateSpider(month); err == nil {
	//	if err = cli.Find(ctx,bson.M{"jyrq":today}).One(&one); err == nil {
	//		return one.Jybz == "1", nil
	//	}
	//}
	return false, err
}