package models

import (
	"fmt"
	"time"
)

type Member struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	FlockId   int       `xorm:"unique(user_id)"`
	UserId    int       `xorm:"unique(flock_id)"`
	CreatedAt time.Time `xorm:"created"`
}

func (this *Member) Add() error {
	return Insert(this)
}

func (this *Member) Del() error {
	return Delete(this)
}

func GetFlocks(userId int) (r []int, err error) {
	engine, err := GetEngine()
	if err != nil {
		return r, err
	}
	defer engine.Close()
	err = engine.Table("member").Cols("flock_id").Find(&r)
	if err != nil {
		return
	}
	return
}

func GetMemberUsers(flockId int) (r []int, err error) {
	engine, err := GetEngine()
	if err != nil {
		return r, err
	}
	defer engine.Close()

	m := make([]*Member, 0)
	err = engine.Table("member").Where("flock_id=?", flockId).Cols("user_id").Find(&m)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	for _, item := range m {
		r = append(r, item.UserId)
	}
	return
}
