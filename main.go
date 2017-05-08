package main

import (
	"github.com/astaxie/beego"
	"hi/models"
	_ "hi/routers"
)

func main() {
	models.SyncModels()
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
