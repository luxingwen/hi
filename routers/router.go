package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"hi/controllers"
)

func init() {

	beego.Router("/ws/join", &controllers.WebSocketControllers{}, "get:WebScoketJoin")
	beego.Router("/webrtc", &controllers.WebRctController{}, "*:Mains")
	beego.Router("/webrtc/ws", &controllers.WebRctController{}, "get:WebRtcJoin")
	beego.Router("/", &controllers.AppController{}, "get:Demo")

	beego.Router("/c/d", &controllers.BaseControlerr{}, "get:ChatDEMO")

	beego.Router("/login", &controllers.UserControler{}, "get:LoginView;post:Login")
	beego.Router("/register", &controllers.UserControler{}, "get:RegisterView;post:Register")
	beego.Handler("/captcha/*", captcha.Server(captcha.StdWidth, captcha.StdHeight))

	apiNs := beego.NewNamespace("/api",
		beego.NSRouter("/user/list", &controllers.UserInFoController{}, "get:UserList"),
		beego.NSRouter("/group", &controllers.GroupController{}, "post:Add;put:Del;get:GetMyAllUser"),
		beego.NSRouter("/friend", &controllers.FriendController{}, "post:Add;put:Del"),
		beego.NSRouter("/flock", &controllers.FlockController{}, "post:Add;put:Del"),
		beego.NSRouter("/member", &controllers.MemberController{}, "post:Add;put:Del"),
		beego.NSRouter("/member/:id([0-9]+)", &controllers.MemberController{}, "get:ListMember"),
		beego.NSRouter("/chat/init", &controllers.ChatController{}, "get:InitInfos"),
		beego.NSRouter("/files", &controllers.FileController{}),
		beego.NSRouter("/skins", &controllers.FileController{}, "get:Skins"),
	)
	beego.AddNamespace(apiNs)
}
