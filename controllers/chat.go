package controllers

import (
	"fmt"
	"hi/models"
	"hi/utils"
)

type ChatController struct {
	AppController
}

type Friends struct {
	Groupname string         `json:"groupname"`
	Id        int            `json:"id"`
	List      []*models.User `json:"list"`
}

type InitInfo struct {
	Mine   *models.User `json:"mine"`
	Friend []*Friends   `json:"friend"`
}

func (this *ChatController) InitInfos() {
	resInfo := new(InitInfo)
	if v, ok := this.GetSession("user").(*models.User); ok {
		resInfo.Mine = v
	}
	friends, err := models.GetFriends(resInfo.Mine.Id)
	if err != nil {
		this.Fail(2, err.Error())
	}
	var groupIds []int
	mGroups := make(map[int][]int)
	for _, v := range friends {
		groupIds = append(groupIds, v.GroupId)
		mGroups[v.GroupId] = append(mGroups[v.GroupId], v.FriendId)
	}
	mGroup, err := models.GetGroups(utils.IdsToCommon(groupIds))
	if err != nil {
		this.Fail(2, err.Error())
	}

	for k, v := range mGroups {
		if group, ok := mGroup[k]; ok {
			friend := &Friends{Groupname: group.GroupName, Id: group.Id}
			users, err := models.GetUsers(v)
			if err != nil {
				this.Fail(2, err.Error())
			}
			for _, uv := range users {
				fmt.Println("users", uv)
			}

			friend.List = users
			resInfo.Friend = append(resInfo.Friend, friend)
		}
	}
	this.Succuess(resInfo)
}
