package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var (
	AppConfig config.Configer
	SqlServer string
	Dev       bool
)

func init() {
	AppConfig = beego.AppConfig
	SqlServer = AppConfig.DefaultString("mysql::server", "")
	Dev = AppConfig.DefaultBool("dev", true)
}

func SyncModels() {
	engine, err := GetEngine()
	if err != nil {
		log.Fatal("get engine err:", err)
	}
	defer engine.Close()
	err = engine.Sync(new(User), new(Friends), new(Group), new(Flock), new(Member))
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(obj interface{}) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	_, err = engine.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

func Update(obj interface{}) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	affected, err := engine.Update(obj)
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("no update.")
	}
	return nil
}

func Delete(obj interface{}) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	_, err = engine.Unscoped().Delete(obj)
	if err != nil {
		return err
	}
	return nil
}

func GetEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", SqlServer)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	return engine, nil
}
