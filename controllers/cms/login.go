package cms

import (
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"test_blob/models"
	"test_blob/utils"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get()  {
	l.TplName="cms/login.html"
}


func (l *LoginController) Post()  {
	username:=l.GetString("username") //可以进阶进行特殊字符过滤操作
	password:=l.GetString("password") //长度检验等操作

	//加密操作
    md5_pwd :=utils.GetMd5(password)

    o:=orm.NewOrm()
    exist :=o.QueryTable(new(models.User)).Filter("user_name",username).Filter("password",md5_pwd).Exist()

	if  exist{
		//跳转页面
		//l.Ctx.WriteString("登录成功....")
		l.SetSession("cms_user_name",username)
		l.Redirect(beego.URLFor("MainController.Get"),302)


	}else {
		l.Redirect(beego.URLFor("LoginController.Get"),302)
	}

}