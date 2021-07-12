package front

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"test_blob/utils"
	"github.com/astaxie/beego/orm"
	"test_blob/models"
)

type RegisterController struct {
beego.Controller
}

func (r *RegisterController) Get()  {
	r.TplName = "front/register.html"

}

func (r *RegisterController) Post()  {

	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")

	if password != repassword {
		r.Ctx.WriteString("两次密码不一致")
	}

	md5_password := utils.GetMd5(password)


	o := orm.NewOrm()
//存入数据库
	user := models.User{
		UserName:username,
		Password:md5_password,
		IsAdmin:2,
		Cover:"static/upload/bq3.png",
	}
	_,err := o.Insert(&user)

	if err != nil {
		fmt.Println(err)
		//r.Ctx.WriteString("用户名已被使用")
	}
   //跳转到/login页面
	r.Redirect("/login",302)


}
