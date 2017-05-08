
[TOC]

# 基于HTML5的聊天系统设计实现

##  摘要   
随着互联网的快速发展，人们的通讯方式发生了巨大变化。网络聊天应用程序自问世以来，由于其实时性，受到了网民的青梅。比较其他互联网通信方式，如电子邮件，您必须等待收件人检查他或她的电子邮件和发送回复，即时通讯的网络聊天工具是在线且实时的。传统的网络聊天应用程序（IM）必须先安装特定的软件和做相应的配置才能使用，此外，它们依赖于平台。对于特别是在公共电脑（如学校图书馆，计算机实验室和网吧）中使用该应用程序的用户而言，这些构成严重问题，大多数情况下，该应用程序大部分未安装。基于Web的即时通讯是这些问题的解决方案，因为它不需要任何下载，安装或配置，而且它是平台无关的。该项目将使用最新的Golang和HTML5技术相结合尝试实现一个基于Web的即时消息应用程序。


关键字：网络聊天、即时消息、IM



## 1、引言

### 1.1 课题背景

随着互联网的快速发展，HTML5逐渐成为了当代互联网web的标准，网络聊天工具作为一种重要的信息交流工具，受到越来越多的网民的喜爱。目前，无论国内外，都出现了许多不错的网络聊天工具，其中国外应用比较广泛的有Facebook、环聊、Skype，国内的有腾讯老大的QQ和微信、啊里巴巴的阿里旺旺。无论个人还是集团等组织机构，对于沟通的需求也在不断发展，随着软件、网络和通信三大信息现代技术的发展，在沟通协作方面有着更多方便、快捷、事实等优势的即时通讯，成为了继电话、email之后有一个完全融入每个人生活的互联网工具。

最近，围绕HTML5 Web Sockets发表了大量热议，其中定义了一个通过Web上的单个套接字运行的全双工通信通道。HTML5 Web套接字不仅仅是对传统HTTP通信的另一种增量增强; 它代表了巨大的进步，特别是对于实时的，事件驱动的web应用程序。


### 1.2 国内外研究现状

目前国内外聊天工具中层出不穷，国外有Facebook、环聊、Skype。在国内，老大当属腾讯的QQ、微信这两款聊天工具，其中这两款聊天工具均有web版本。而且微信的web版本充分利用了HTML5的新特性，使得在浏览器中体验也不亚于客户端。

许多网站为了给客户带来更好的服务，纷纷在网站上添加了客户功能。例如：淘宝、天猫、京东的客户服务，还有三大营运商的客服服务。均在web上实现了即时通讯。




### 1.3 课题研究内容

本次课题研究为使用Go语言和HTML5实现一个web即时聊天系统。系统主要的模块有登录、注册模块，一对一聊天，群聊聊天，添加好友，添加群组等功能

### 1.4 课题研究方法
本课题的研究方法为理论+实践。
国内外搜集HTML5与Golang相关的学习资料。
  使用go语言作为后端开发语言，HTML5做为前端开发语言，sublime text3开发工具。使用mysql数据库，




## 2. HTML5实时应用程序相关技术简介

### 2.1 实时Web应用程序的简史
Web是围绕这样的想法构建的，客户端的工作是从服务器请求数据，服务器的工作是满足这些请求。这种模式多年来没有被挑战，但随着2005年左右引入AJAX，许多人开始探索在客户端和服务器之间建立双向连接的可能性。

Web应用程序已经长大了，现在比以前消耗更多的数据。阻止他们回来的最大的事情是客户端发起的事务的传统HTTP模型。为了克服这一点，设计了许多不同的策略以允许服务器将数据推送到客户端。这些策略中最流行的一种是长期轮询。这包括保持HTTP连接打开，直到服务器有一些数据推送到客户端。

轮询，就是你去车站上厕所。车站有很多人，你第一次去，发现厕所有人，于是你又回来等，过了几分钟，再去，发现还是有人。然后在回来等待，然后又去。以此往复。

所有这些解决方案的问题是它们携带HTTP的开销。每次发出HTTP请求时，一大堆标头和Cookie数据都会传输到服务器。这可能累加需要传输的相当大量的数据，这反过来增加延迟。如果你正在构建类似于基于浏览器的游戏，减少延迟是保持事务顺利运行的关键。最糟糕的部分是，很多这些头和cookie实际上不需要满足客户端的请求。

### 2.2 websocket是什么

Websocket是html5提出的一个协议规范，参考rfc6455。
websocket约定了一个通信的规范，通过一个握手的机制，客户端（浏览器）和服务器（webserver）之间能建立一个类似tcp的连接，从而方便c－s之间的通信。在websocket出现之前，web交互一般是基于http协议的短连接或者长连接。

WebSocket是为解决客户端与服务端实时通信而产生的技术。websocket协议本质上是一个基于tcp的协议，是先通过HTTP/HTTPS协议发起一条特殊的http请求进行握手后创建一个用于交换数据的TCP连接，此后服务端与客户端通过此TCP连接进行实时通信。


### 2.3 WebSockets如何工作

WebSockets在客户端和服务器之间提供持久连接，双方可以随时使用它们来开始发送数据。客户端通过称为WebSocket握手的过程建立WebSocket连接。此过程从客户端向服务器发送常规HTTP请求开始。Upgrade此请求中包含一个标头，通知服务器客户端希望建立WebSocket连接。

以下是初始请求标头的简化示例。

GET ws://websocket.example.com/ HTTP/1.1
Origin: http://example.com
Connection: Upgrade
Host: websocket.example.com
Upgrade: websocket

WebSocket URL使用ws方案。还有wss对等的安全WebSocket连接HTTPS。
如果服务器支持WebSocket协议，则它同意升级并通过Upgrade响应中的标头进行通信。

HTTP/1.1 101 WebSocket Protocol Handshake
Date: Wed, 16 Oct 2013 10:07:34 GMT
Connection: Upgrade
Upgrade: WebSocket

现在握手已完成，初始HTTP连接将替换为使用相同底层TCP / IP连接的WebSocket连接。此时，任一方都可以开始发送数据。

使用WebSockets，您可以传输尽可能多的数据，而不会产生与传统HTTP请求相关的开销。数据通过WebSocket作为消息传输，每个消息由一个或多个包含您要发送的数据（有效载荷）的帧组成。为了确保消息可以正确重建，当它到达客户端时，每个帧都有4-12字节的有关有效负载的数据前缀。使用这种基于帧的消息系统有助于减少传输的非有效载荷数据量，从而显着减少延迟。

 



## 3、Go语言简介

Golang编程语言是一个开源项目，是简洁的，使程序员高效。它的并发机制使得编写高并发的应用程序更容易，而其新颖的类型系统实现了灵活的模块化程序构造。Go快速编译到机器代码，但具有垃圾收集的方便和运行时反射的力量。它是一个快速，静态类型，编译的语言，感觉像一个动态类型，解释语言。

Go非常适合编写高性能，并发服务器应用程序和工具，Go对网络编程提供了全面的支持，现在很难找到一个没有上网功能的小工具，使用go编写网络程序将变得很容易。

在诸如C＃，Java或Node.js之类的语言中，需要采用复杂的线程代码和巧妙使用锁来保持所有客户端的同步。在这个项目中，我们将看到，Go通过其特有的channel帮助我们。



## 4、开发环境及配置

### 4.1 系统开发环境、及开发工具
系统开发环境：macOS Sierra 10.12.4 、Go version go1.8 darwin/amd64、mysql  Ver 14.14 Distrib 5.7.17, for osx10.12 (x86_64) using  EditLine wrapper
开发工具： sublime Text3

### 4.2系统运行环境
服务器端程序最终可编译生成对应平台的可执行文件且可运行在各个平台上。本文主要在macos 上运行以及测试。
客户端程序主要运行在Google Chrome和火狐浏览器上。

### 4.3系统开发引用第三方框架以及包
服务器程序主要依赖beego web开发框架。
服务器引入的第三方包包括：

 - "github.com/dchest/captcha"
 - "github.com/gorilla/websocket"
  
客户端基于Bootstrap框架。


## 5、项目需求分析与数据库设计

### 5.1 需求分析

#### 5.1.1 登录注册功能

每个人可以输入自己的邮箱、用户名和密码进行注册，后端录入数据库，如果邮箱在数据库已经存在，则提示该用户已存在。 用户通过输入邮箱和正确的密码登录到聊天系统，如果没有匹配到用户，则提示用户名或者密码错误。

#### 5.1.2 查找用户好友功能

用户可以通过好友查找输入框来查找自己已经存在的好友。

#### 5.1.3 添加好友功能

用户可以查找系统存在的用户，并可以选择将其添加好友，此时会给对方发送一个添加好友的请求。对方同意或者拒接后，用户自己将收到对方的消息，如果是同意，则可选择分组将对方添加到自己的好友列表里。反之是拒接则提示对方拒接添加好友等消息。

#### 5.1.4 用户聊天

用户可以选择好友或者群组打开聊天面板，通过内容输入框输入聊天内容，点击发送，将消息发送给对方。

### 5.2 数据库设计

系统需要的数据表存放在mysql数据库中。需要一个用户user表来存放用户数据。群组表存放群组信息，分组表存放用户的分组信息，朋友表存放用户的朋友信息。以下为表属性的详细介绍。

#### 5.2.1 用户表：user

列名 | 属性类型 | 属性作用
------- | -------|-----
 id| int |id
email|string|用户邮箱
user_name|sring|用户名
pass_word|string|密码
nick_name|string|用户昵称
status|string|用户状态
sign|string|用户签名
sex|int|用户性别
avatar|string|用户头像
created|date|创建时间
updated|date|更新时间

#### 5.2.2 群组表：flock

列名 | 属性类型 | 属性作用
------- | -------|-----
id|int|自增id
name|string|群名称
owner|int|拥有者（创建者）
avatar|string|群头像
created|date|创建时间
updated|date|更新时间

#### 5.2.3 分组表：group

列名 | 属性类型 | 属性作用
------- | -------|-----
id|int|自增id
group_name|string|分组名称
user_id|int|用户id
created|date|创建时间
updated|date|更新时间


#### 5.2.4 朋友表：friends

列名 | 属性类型 | 属性作用
------- | -------|-----
id|int|自增id
group_id|int|组id
friend_id|int|朋友id
user_id|int|用户id
created|date|创建时间
updated|date|更新时间

## 6、系统详细设计与实现

### 6.1 程序功能分析与设计

用户首先通过登录／注册进入到聊天系统，然后浏览器客户端向服务器发送Websocket连接请求，服务器端成功解析客户端发出的websocket连接请求后并返回应答信息给客户端，然后客户端和服务器建立起了websocket连接通道。用户与用户聊天时，A发送给B的消息，服务器立即推送给B。在群组里聊天时，服务器支持广播功能，每一个用户发的聊天消息都会实时推送到群组里面的所有用户，当用户退出时。服务端要断开相应的连接，并删除相应的连接信息。
下面我们开始来设计与实现该系统，我们的服务器使用golang语言实现，客户端使用HTML5编写，从而运行在浏览器内。

- 系统活动图

下图为系统的活动图，它描述了用户与系统交互的过程。

```flow
st=>start: start:>http://127.0.0.1:8877[blank]
e=>end:>
register=>operation: 注册
login=>operation: 登录
addContact=>condition: 添加新朋友？
addFriend=>operation: 添加新朋友
selectContact=>operation: 选择联系人或群组
startChat=>operation: 开始聊天
logout=>operation: 退出
cond=>condition: 登录
or 注册?

st->cond
cond(yes)->login->addContact(yes)->addFriend(right)->selectContact
cond(no)->register(right)->addContact(no)->selectContact->startChat->logout->e
```
	

- 用例图

下图为系统的用例图。它显示显示用户如何与系统进行交互。
![](http://bmob-cdn-9735.b0.upaiyun.com/2017/04/15/c5cab5124096791480d41bf4236eeaf7.png)


- 事件图
下图为事件图，它概述了用户使用系统的过程/步骤。
![](http://bmob-cdn-9735.b0.upaiyun.com/2017/04/15/690ca7aa401841ba8058a11916cdec11.png)



### 6.2用户Websocket连接实现
客户端
```js
 var socket=new WebSocket("ws://"+window.location.host+"/ws/join");
```

服务器

```go
ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Println("websocket connection err:", err)
		this.Redirect("/", 302)
		return
	}
```

### 6.3登录功能实现

登录输入只需要输入正确的用户账户以及密码就能登录，从登录页面输入用户名和密码，然后提交表单。服务器端根据用户提交的密码，计算出md5 字符串，由此去数据库获取用户的信息，并把该用户对象返回，保存在session中。

以下为服务器端登录功能主要代码
```go
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
```

登录界面展示

![](http://bmob-cdn-9735.b0.upaiyun.com/2017/04/05/da6509fa407cfb1f804d03dae1d45cd6.png)


### 6.4 注册功能实现

用户注册统一为邮箱注册，注册时候需要输入用户名、用户密码、用户邮箱等基本信息。由前端HTML5页面提交表单到服务器端，服务器端主要根据用户提交的密码计算出相应的md5，然后把相关的信息插入到数据库，如果成功插入，把此对象保存到session中去。

以下为服务器端注册功能主要代码
```go
func (this *User) Add() error {
	return Insert(this)
}

func (this *UserControler) Register() {
	username := this.GetString("username", "")
	password := this.GetString("password", "")
	email := this.GetString("email", "")
	sex, _ := this.GetInt("sex", 1)

	user := &models.User{UserName: username, PassWord: password, Email: email, Sex: sex}
	if user.UserName == "" || user.PassWord == "" || user.Email == "" {
		this.Fail(2, "用户名或者密码或者邮箱为空")
	}
	err := user.Add()
	if err != nil {
		this.Fail(2, "register err.")
	}
	this.SetSession("user", user)
	this.Redirect("/", 301)
}

```

注册页面展示

![](http://bmob-cdn-9735.b0.upaiyun.com/2017/04/05/83dd41c7403ef18580e90946ca857ad5.png)

###  6.5主页面实现

当用户登录／注册成功后。会进入到主页面，这时候前端会发送一个WebSocket连接请求，当成功以服务器建立起连接后，在请求一个初始化信息接口，从而获得服务器返回的 用户本身的信息以及用户的好友信息列表和用户的群组信息列表。

以下为服务器端主页面实现核心代码
```go
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
			friend.List = users
			resInfo.Friend = append(resInfo.Friend, friend)
		}
	}
	this.Succuess(resInfo)
}
```

主页面展示

![](http://bmob-cdn-9735.b0.upaiyun.com/2017/04/15/8eb28e1240f476d780a4d3b877a4e88c.png)



### 6.6聊天功能实现

当用户发送消息的时候。直接往与服务器端建立的WebSocket连接里写入信息，服务器端从该连接通道里面读取出消息内容。解析成消息体，然后加入到消息队列中去。最后在广播，广播即为推送消息到用户端。

其中用户发送的消息体如下：
```go
type Event struct {
	Type string `json:"type"`//消息类型
	Data struct {// 消息体
		Mine struct {////包含我发送的消息及我的信息
			Avatar   string `json:"avatar"`
			Id       int    `json:"id"`
			Mine     bool   `json:"mine"`
			Username string `json:"username"`
			Content  string `json:"content"`
		} `json:"mine"`
		To struct {//对方的信息
			Avatar   string `json:"avatar"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Id       int    `json:"id"`
			Type     string `json:"type"`
		} `json:"to"`
	} `json:"data"`
}
```

其中服务器端从WebSocket连接读取消息并加入到消息队列的具体实现如下：

```go
//从websocket连接读取消息
		_, b, err := ws.ReadMessage()
		if err != nil {
			return
		}
		//解析成消息体对象
		var event *models.Event
		if err = json.Unmarshal(b, &event); err != nil {
			fmt.Println(err)
			continue
		}
		//加入到消息队列
		chatroom.publish <- event
```

最后广播消息到需要推送的客户端具体实现如下

```go

func broadcastWebSocket(event *models.Event, msg *models.Message) {
	if event.Data.To.Type == "friend" {
		msg.Id = strconv.Itoa(event.Data.Mine.Id)
		uId := event.Data.To.Id
		if item, ok := chatroom.MSubscriber[uId]; ok {
			if err := item.Conn.WriteJSON(msg); err != nil {
				fmt.Println("发送失败：", err)
				chatroom.remove <- item
			}
		}
	} else if event.Data.To.Type == "group" {
		msg.Id = strconv.Itoa(event.Data.To.Id)
		groupId := event.Data.To.Id
		uIds, err := models.GetMemberUsers(groupId)
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
```



聊天界面展示

![](http://bmob-cdn-9735.b0.upaiyun.com/2017/04/15/92d674fd408739f480b85c0da852fbd9.png)



### 6.7 添加好友功能实现

添加好友面板首先会列出该系统内目前在线的玩家列表。用户在搜索框内输入用户名或者用户邮箱相关信息，服务器会进行模糊查询并把相关的用户列表信息返回。用户点击添加好友按钮后会给对方发送一条添加好友消息请求。对方同意后会把

## 7、运行测试




## 参考文献



