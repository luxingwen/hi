package controllers

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
	AppController
}

func (this *FileController) Post() {
	req := this.Ctx.Request
	f, h, err := req.FormFile("file")
	if err != nil {
		this.Fail(2, err.Error())
	}
	upload := "upload/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day())
	err = os.MkdirAll(upload, 0777)
	if err != nil {
		this.Fail(2, err.Error())
	}
	fs, err := os.Create(upload + "/" + h.Filename)
	io.Copy(fs, f)
	m := map[string]string{"src": "/" + upload + "/" + h.Filename}
	this.Succuess(m)
}

func (this *FileController) Skins() {
	dir := "static/skin"
	var pic []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		this.Fail(404, "不能读取文件")
	}
	for _, item := range files {
		if !strings.HasSuffix(item.Name(), ".jpg") && !strings.HasSuffix(item.Name(), ".png") {
			continue
		}
		pic = append(pic, "/"+dir+"/"+item.Name())
	}
	this.Succuess(pic)
}
