package controllers

import (
	"github.com/astaxie/beego"
)

type BaseControlerr struct {
	beego.Controller
}

func (this *BaseControlerr) Mains() {
	this.TplName = "index.html"
}

func (this *BaseControlerr) ChatDEMO() {
	this.TplName = "demo/index.html"
}
func (this *BaseControlerr) Succuess(data interface{}) {
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": data}
	this.ServeJSON()
}

func (this *BaseControlerr) Fail(code int, msg string) {
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
	this.StopRun()
}
