package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hi/models"
)

type UserInFoController struct {
	AppController
}

func (this *UserInFoController) UserList() {
	users, err := models.UserList()
	if err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(users)
}

func (this *UserInFoController) OnlineList() {
	//query := this.GetString("q")

}

func (this *UserControler) Update() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	var user *models.User
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("register func, json unmarshal err:%v", err)
		return
	}
	if err := user.Update(); err != nil {
		fmt.Println("update error. error: ", err)
		return
	}
	this.Data["json"] = map[string]interface{}{"msg": "success.", "code": 200}
	this.ServeJSON()
}
