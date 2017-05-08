package controllers

import (
	"fmt"
	"github.com/dchest/captcha"
	"hi/models"
)

type UserControler struct {
	BaseControlerr
}

func (this *UserControler) LoginView() {
	this.Data["CaptchaId"] = captcha.New()
	this.TplName = "login.html"
}

func (this *UserControler) Login() {
	// body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	// var user struct {
	// 	UserName string `json:"username"`
	// 	PassWord string `json:"password"`
	// 	Email    string `json:"email"`
	// }
	// err := json.Unmarshal(body, &user)
	// if err != nil {
	// 	fmt.Println("Login func: json unmar shal err:%v", err)
	// 	this.Fail(1, "login failer.")
	// }
	username := this.GetString("username", "")
	password := this.GetString("password", "")
	rUser, err := models.Login(username, "", password)
	if err != nil {
		fmt.Println("err: ", err)
		this.Fail(1, "login failer.")
	}

	this.SetSession("user", rUser)
	this.Redirect("/", 301)
	//this.Succuess(rUser)
}

func (this *UserControler) RegisterView() {
	this.Data["CaptchaId"] = captcha.New()
	this.TplName = "register.html"
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
