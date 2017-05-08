package models

import (
	"errors"
	"time"
)

type User struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	Email     string    `xorm:"varchar(30) unique" json:"email"`
	UserName  string    `xorm:"varchar(20) unique" json:"username"`
	PassWord  string    `json:"password,omitempty"`
	NickName  string    `json:"nickName"`
	Status    string    `json:"status"`
	Sign      string    `json:"sign"`
	Sex       int       `json:"sex"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `xorm:"created"`
	UpdateAt  time.Time `xorm:"updated"`
}

func (this *User) Add() error {
	return Insert(this)
}

func (this *User) Del() error {
	return Delete(this)
}

func (this *User) Update() error {

	return Update(this)
}

func GetUser(id int) (r *User, err error) {
	engin, err := GetEngine()
	if err != nil {
		return nil, err
	}
	defer engin.Close()

	r = new(User)
	has, err := engin.Where("id=?", id).Get(r)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("not found user.")
	}
	r.PassWord = ""
	return
}

func GetUsers(ids []int) (r []*User, err error) {
	engin, err := GetEngine()
	if err != nil {
		return nil, err
	}
	defer engin.Close()
	err = engin.In("id", ids).Cols("id", "user_name", "nick_name", "avatar", "status", "sign").Find(&r)
	if err != nil {
		return
	}
	return
}

func UserList() (r []*User, err error) {
	engin, err := GetEngine()
	if err != nil {
		return nil, err
	}
	defer engin.Close()
	err = engin.Cols("id", "user_name", "nick_name", "avatar", "status", "sign").Find(&r)
	return
}

func Login(userName, email, pwd string) (user *User, err error) {
	engin, err := GetEngine()
	if err != nil {
		return nil, err
	}
	defer engin.Close()
	user = new(User)
	has, err := engin.Where("(user_name=? OR email=?) AND pass_word=?", userName, email, pwd).Get(user)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("not found user.")
	}
	user.PassWord = ""
	return
}
