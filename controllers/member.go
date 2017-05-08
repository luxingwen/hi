package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hi/models"
	"strconv"
)

type MemberController struct {
	AppController
}

func (this *MemberController) Add() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	member := new(models.Member)
	err := json.Unmarshal(body, &member)
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
	member.UserId = user.Id
	if member.FlockId == 0 {
		this.Fail(2, "no flock")
	}
	if err = member.Add(); err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(member)
}

func (this *MemberController) Del() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	member := new(models.Member)
	err := json.Unmarshal(body, &member)
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
	member.UserId = user.Id
	if member.FlockId == 0 {
		this.Fail(2, "no flock")
	}
	err = member.Del()
	if err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(nil)
}

func (this *MemberController) ListMember() {
	fmt.Println("list member:")
	flockIdStr := this.Ctx.Input.Param(":id")
	flockId, _ := strconv.Atoi(flockIdStr)
	userIds, err := models.GetMemberUsers(flockId)
	if err != nil {
		this.Fail(2, err.Error())
	}
	fmt.Println("user ids: ", userIds)
	users, err := models.GetUsers(userIds)
	if err != nil {
		this.Fail(2, err.Error())
	}
	m := map[string]interface{}{"list": users}
	this.Succuess(m)
}
