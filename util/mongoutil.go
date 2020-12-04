package util

import(
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/qiniu/qmgo"
)


var (
	host = beego.AppConfig.String("host")
	dbName = beego.AppConfig.String("database")
	username = beego.AppConfig.String("username")
	passwd = beego.AppConfig.String("password")
)

//连接db,
func ConnectWithColl(coll string) (context.Context, *qmgo.QmgoClient)  {
	ctx := context.Background()
	mongodb := fmt.Sprintf("mongodb://%s:%s@%s", username, passwd, host)
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: mongodb, Database: dbName, Coll: coll})
	if err != nil {
		panic(err)
	}
	return ctx, cli
}



