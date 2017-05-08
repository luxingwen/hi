package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hi/models"
)

type FriendController struct {
	AppController
}

func (this *FriendController) Add() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	friend := new(models.Friends)
	err := json.Unmarshal(body, &friend)
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
	friend.UserId = user.Id
	if friend.GroupId == 0 || friend.FriendId == 0 {
		this.Fail(2, "jj")
	}
	if err = friend.Add(); err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(friend)
}

func (this *FriendController) Del() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	friend := new(models.Friends)
	err := json.Unmarshal(body, &friend)
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
	friend.UserId = user.Id
	if friend.FriendId == 0 {
		this.Fail(2, "jj")
	}
	if err = friend.Del(); err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(friend)
}
