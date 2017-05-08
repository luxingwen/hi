package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type WebRctController struct {
	BaseControlerr
}

func (this *WebRctController) Mains() {
	this.TplName = "webrtc.html"
}

type WebRtc struct {
	EventName string `json:"eventName"`
	Data      struct {
		Room       string      `json:"room"`
		SocketId   int         `json:"socketId"`
		MeSocketId int         `json:"meSocketId"`
		Label      interface{} `json:"label"`
		Candidate  interface{} `json:"candidate"`
		Sdp        interface{} `json:"sdp"`
	} `json:"data"`
}

var count int = 0

type UWebRtc struct {
	Id   int
	Conn *websocket.Conn
}

var us []*UWebRtc

func (this *WebRctController) WebRtcJoin() {
	// if this.User == nil {
	// 	if v, ok := this.GetSession("user").(*models.User); ok {
	// 		this.User = v
	// 	} else {
	// 		this.Redirect("/", 302)
	// 		return
	// 	}
	// }
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Println("websocket connection err:", err)
		this.Redirect("/", 302)
		return
	}
	for {
		_, b, err := ws.ReadMessage()
		if err != nil {
			return
		}
		var data WebRtc
		if err := json.Unmarshal(b, &data); err != nil {
			fmt.Println("unmar shal err: ", err)
			return
		}
		if data.EventName == "__join" {
			fmt.Println("__join")
			count++
			u := &UWebRtc{Id: count, Conn: ws}
			JoinWebRtc(u)
			var ids []int
			for i := 1; i < count; i++ {
				ids = append(ids, i)
			}
			rsp := map[string]interface{}{
				"eventName": "_peers",
				"data": map[string]interface{}{
					"connections": ids,
					"you":         count,
				},
			}
			ws.WriteJSON(rsp)
		}

		if data.EventName == "__ice_candidate" {
			for _, item := range us {
				if item.Id == data.Data.SocketId {
					rsp := map[string]interface{}{
						"eventName": "_ice_candidate",
						"data": map[string]interface{}{
							"label":     data.Data.Label,
							"candidate": data.Data.Candidate,
							"socketId":  data.Data.MeSocketId,
						},
					}
					item.Conn.WriteJSON(rsp)
					fmt.Println("__ice_candidate write: ", rsp)
				}
			}
		}

		if data.EventName == "__offer" {
			for _, item := range us {
				if item.Id == data.Data.SocketId {
					rsp := map[string]interface{}{
						"eventName": "_offer",
						"data": map[string]interface{}{
							"sdp":      data.Data.Sdp,
							"socketId": data.Data.MeSocketId,
						},
					}
					item.Conn.WriteJSON(rsp)
					fmt.Println("__offer write: ", rsp)
				}
			}
		}

		if data.EventName == "__answer" {
			for _, item := range us {
				if item.Id == data.Data.SocketId {
					rsp := map[string]interface{}{
						"eventName": "_answer",
						"data": map[string]interface{}{
							"sdp":      data.Data.Sdp,
							"socketId": data.Data.MeSocketId,
						},
					}
					item.Conn.WriteJSON(rsp)
					fmt.Println("__answer write: ", rsp)
				}
			}
		}
		fmt.Println("count: ", count)
	}

}

func JoinWebRtc(u *UWebRtc) {
	for _, item := range us {
		rsp := map[string]interface{}{
			"eventName": "_new_peer",
			"data": map[string]int{
				"socketId": u.Id,
			},
		}
		item.Conn.WriteJSON(rsp)
		fmt.Println("rsp: ", rsp)
	}
	us = append(us, u)
}
