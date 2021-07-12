package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "test_blob/routers"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_"test_blob/models"
	"test_blob/utils"
)
//连接数据库
func init()  {
	username,_:=beego.AppConfig.String("username")

	password,_:= beego.AppConfig.String("password")

	host,_:=beego.AppConfig.String("host")

	port,_:=beego.AppConfig.String("port")

	database,_:=beego.AppConfig.String("database")
	//拼接字符串
	datasours :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",username,password,host,port,database)
	//驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//必须要有一个default,连接数据库
	orm.RegisterDataBase("default","mysql",datasours)

	fmt.Println("连接数据库")
	//自动连接数据库
	name :="default"
	force :=false
	verbose :=true
	err :=orm.RunSyncdb(name,force,verbose)
	if err != nil {
		panic(err)
	}
}

func main() {
	//过滤登录
	beego.InsertFilter("cms/main/*",beego.BeforeRouter,utils.CmsLoginFilter)
	//beego.InsertFilter("/comment",beego.BeforeRouter,utils.FrontLoginFilter)
	orm.RunCommand()
	beego.Run()
}

