package models

import (
	"time"
)

type Flock struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	Name      string    `json:"name"`
	Owner     int       `json:"owner"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *Flock) Add() error {
	return Insert(this)
}

func (this *Flock) Del() error {
	return Delete(this)
}
