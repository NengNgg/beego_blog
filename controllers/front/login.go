package front

import (
	beego "github.com/beego/beego/v2/server/web"
	"test_blob/utils"
	"github.com/astaxie/beego/orm"
	"test_blob/models"

)

type FrontLoginController struct {
beego.Controller
}

func (l *FrontLoginController) Get()  {
	l.TplName = "front/login.html"

}

func (l *FrontLoginController) Post()  {

	username := l.GetString("username")  // 特殊字符过滤。?/
	password := l.GetString("password") // 长度校验

	md5_pwd := utils.GetMd5(password)

	o := orm.NewOrm()

	exist := o.QueryTable(new(models.User)).Filter("user_name",username).Filter("password",md5_pwd).Exist()

	if exist {

		l.SetSession("front_user_name",username)

		l.Redirect(beego.URLFor("IndexController.Get"),302)
	}else {
		//这里没有注册或者密码错误直接跳转到本页面
		l.Redirect(beego.URLFor("FrontLoginController.Get"),302)
	}

}
