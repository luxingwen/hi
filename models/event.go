package models

/*

{"type":"chatMessage","data":{"mine":{"username":"ss1","avatar":"http://a4.topitme.com/o129/1012939423c49a2a0b.jpg","id":2,"mine":true,"content":"123"},"to":{"id":1,"email":"","username":"ss2","nickName":"ss2","status":"","sign":"","avatar":"http://a4.topitme.com/l/201010/01/12859073846243.jpg","CreatedAt":"0001-01-01T00:00:00Z","UpdateAt":"0001-01-01T00:00:00Z","name":"ss2","type":"friend"}}}

*/

type Event struct {
	Type string `json:"type"`
	Data struct {
		Mine struct {
			Avatar   string `json:"avatar"`
			Id       int    `json:"id"`
			Mine     bool   `json:"mine"`
			Username string `json:"username"`
			Content  string `json:"content"`
		} `json:"mine"`
		To struct {
			Avatar   string `json:"avatar"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Id       int    `json:"id"`
			Type     string `json:"type"`
		} `json:"to"`
	} `json:"data"`
}

type Message struct {
	UserName  string `json:"username"`
	Avatar    string `json:"avatar"`
	Id        string `json:"id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Mine      bool   `json:"mine"`
	Fromid    string `json:"fromid"`
	Timestamp int64  `json:"timestamp"`
}
