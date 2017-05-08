package controllers

import (
	"github.com/gorilla/websocket"
	"hi/models"
	"strconv"
	"time"
)

type Subscriber struct {
	User *models.User
	Conn *websocket.Conn
}

type ChatRoom struct {
	MSubscriber map[int]*Subscriber
	add         chan *Subscriber
	remove      chan *Subscriber
	publish     chan *models.Event
	MGroup      map[int][]int
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		MSubscriber: make(map[int]*Subscriber),
		add:         make(chan *Subscriber, 100),
		remove:      make(chan *Subscriber, 100),
		publish:     make(chan *models.Event, 100),
		MGroup:      make(map[int][]int),
	}
}

func (this *ChatRoom) Join(user *models.User, conn *websocket.Conn) {
	this.add <- &Subscriber{User: user, Conn: conn}

}

func Leave(user *models.User) {
	chatroom.remove <- &Subscriber{User: user}
}

func (c *ChatRoom) removeSub(sub *Subscriber) {
	delete(c.MSubscriber, sub.User.Id)
}

var chatroom *ChatRoom

func init() {
	chatroom = NewChatRoom()
	go chatroom.chatroom()
}

func (c *ChatRoom) chatroom() {
	for {
		select {
		case sub := <-c.add:
			c.MSubscriber[sub.User.Id] = sub
		case publish := <-c.publish:
			msg := &models.Message{UserName: publish.Data.Mine.Username, Avatar: publish.Data.Mine.Avatar, Type: publish.Data.To.Type, Content: publish.Data.Mine.Content}
			msg.Mine = false
			msg.Fromid = strconv.Itoa(publish.Data.Mine.Id)
			msg.Timestamp = time.Now().Unix()
			broadcastWebSocket(publish, msg)
		case unsub := <-c.remove:
			c.removeSub(unsub)
		}
	}
}
