package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"hi/models"
	"strconv"
)

type WebSocketControllers struct {
	AppController
}

func (this *WebSocketControllers) WebScoketJoin() {
	var user *models.User
	if v, ok := this.GetSession("user").(*models.User); ok {
		user = v
	} else {
		this.Redirect("/", 302)
		return
	}
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Println("websocket connection err:", err)
		this.Redirect("/", 302)
		return
	}
	chatroom.Join(user, ws)
	defer Leave(user)
	for {
		_, b, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("body: ", string(b))
		var event *models.Event
		if err = json.Unmarshal(b, &event); err != nil {
			fmt.Println(err)
			continue
		}

		chatroom.publish <- event
	}
	return
}

func broadcastWebSocket(event *models.Event, msg *models.Message) {
	fmt.Println("event: ", event)

	if event.Data.To.Type == "friend" {
		msg.Id = strconv.Itoa(event.Data.Mine.Id)
		uId := event.Data.To.Id
		if item, ok := chatroom.MSubscriber[uId]; ok {
			fmt.Println("msg: ", msg)
			fmt.Println("item username: ", item.User.UserName)
			if err := item.Conn.WriteJSON(msg); err != nil {
				fmt.Println("发送失败：", err)
				chatroom.remove <- item
			}
		}
	} else if event.Data.To.Type == "group" {
		msg.Id = strconv.Itoa(event.Data.To.Id)
		groupId := event.Data.To.Id
		uIds, err := models.GetMemberUsers(groupId)
		fmt.Println("msg: ", msg)
		if err != nil {
			fmt.Println("获取群员资料失败：", err)
		} else {
			for id, _ := range uIds {
				if item, ok := chatroom.MSubscriber[id]; ok {
					if err := item.Conn.WriteJSON(msg); err != nil {
						fmt.Println("发送失败：", err)
						chatroom.remove <- item
					}
				}
			}
		}
	}

}
