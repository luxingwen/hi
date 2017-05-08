package models

import (
	"time"
)

type Friends struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	GroupId   int       `json:"groupId"`
	FriendId  int       `json:"friendId"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *Friends) Add() error {
	return Insert(this)
}

func (this *Friends) Del() error {
	return Delete(this)
}

func (this *Friends) Update() error {
	return Update(this)
}

func GetFriends(userId int) (r []*Friends, err error) {
	engine, err := GetEngine()
	if err != nil {
		return nil, err
	}
	err = engine.Where("user_id=?", userId).Find(&r)
	if err != nil {
		return nil, err
	}
	return
}
