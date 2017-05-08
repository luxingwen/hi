package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hi/models"
)

type GroupController struct {
	AppController
}

func (this *GroupController) Add() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	group := new(models.Group)
	err := json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println(err)
		this.Fail(2, err.Error())
	}
	var user *models.User
	if v, ok := this.GetSession("user").(*models.User); ok {
		user = v
	} else {
		this.Fail(2, "invalid user.")
	}
	group.UserId = user.Id
	if group.GroupName == "" {
		this.Fail(2, "empty group name")
	}
	if err = group.Add(); err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(group)
}

type GroupUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (this *GroupController) GetMyAllUser() {
	users, err := models.UserList()
	if err != nil {
		this.Fail(2, err.Error())
	}
	var gs []*GroupUser
	for _, item := range users {
		g := &GroupUser{Id: item.Id, Name: item.UserName}
		gs = append(gs, g)
	}
	this.Succuess(gs)
}

func (this *GroupController) Del() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	group := new(models.Group)
	err := json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println(err)
		this.Fail(2, err.Error())
	}
	var user *models.User
	if v, ok := this.GetSession("user").(*models.User); ok {
		user = v
	} else {
		this.Fail(2, "invalid user.")
	}
	group.UserId = user.Id
	if group.Id == 0 {
		this.Fail(2, "empty group id")
	}
	if err = group.Del(); err != nil {
		this.Fail(2, "del group err")
	}
	this.Succuess(err)
}
