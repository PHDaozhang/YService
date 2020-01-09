package main

import (
	_ "YService/routers"
	"github.com/astaxie/beego"
	"time"
	"tsEngine/tsRand"
)

func main() {



	tsRand.Seed(time.Now().UnixNano())

	beego.Run()
}

