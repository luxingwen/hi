package models

import (
	"time"
)

type Group struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	GroupName string    `json:"groupname"`
	Avatar    string    `json:"avatar"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *Group) Add() error {
	return Insert(this)
}

func (this *Group) Del() error {
	return Delete(this)
}

func (this *Group) Update() error {
	return Update(this)
}

func GetGroups(ids []int) (r map[int]*Group, err error) {
	r = make(map[int]*Group)
	engine, err := GetEngine()
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	err = engine.In("id", ids).Find(&r)
	if err != nil {
		return nil, err
	}
	return
}
