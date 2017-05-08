package controllers

import (
	"github.com/astaxie/beego"
	//"hi/models"
	//"log"
)

type AppController struct {
	beego.Controller
}

func (this *AppController) Prepare() {
	// if models.Dev {
	// 	user, err := models.GetUser(1)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	this.SetSession("user", user)
	// }

	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/login", 302)
	}
}

func (this *AppController) Demo() {
	this.TplName = "chat.html"
}

func (this *AppController) Succuess(data interface{}) {
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": data}
	this.ServeJSON()
}

func (this *AppController) Fail(code int, msg string) {
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
	this.StopRun()
}
