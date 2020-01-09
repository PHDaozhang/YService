package main

import (
	"YService/conf"
	"YService/core/cache"
	_ "YService/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
	"tsEngine/tsDb"
	"os"
	"tsEngine/tsRand"
)

func main() {
	logs.SetLogger("file", `{"filename":"./logs/logs.log", "perm":"0775", "maxDays":15}`)
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)

	err :=  tsDb.ConnectDbFormatConfig("mysql", conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
	if err != nil {
		logs.Error("连接数据库失败,default")
		os.Exit(1)
	}


	orm.RunSyncdb("default",false,true)

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir ["/swagger"] = "swagger"

	cache.InitRedis();

	tsRand.Seed(time.Now().UnixNano())

	beego.Run()
}

