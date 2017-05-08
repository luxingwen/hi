package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hi/models"
)

type FlockController struct {
	AppController
}

func (this *FlockController) Add() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	flock := new(models.Flock)
	err := json.Unmarshal(body, &flock)
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
	flock.Owner = user.Id
	if flock.Name == "" {
		this.Fail(2, "invalid name.")
	}
	if err = flock.Add(); err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(flock)
}

func (this *FlockController) Del() {
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	flock := new(models.Flock)
	err := json.Unmarshal(body, &flock)
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
	flock.Owner = user.Id
	if err = flock.Del(); err != nil {
		this.Fail(2, err.Error())
	}
	member := &models.Member{FlockId: flock.Id}
	if err = member.Del(); err != nil {
		this.Fail(2, err.Error())
	}
	this.Succuess(nil)
}
