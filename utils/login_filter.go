package utils

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func FrontLoginFilter(ctx *context.Context)  {
	//判断session是否为空
	front_user_name :=ctx.Input.Session("cms_user_name")
	if front_user_name==nil {
		//为空的话就跳转到登录页面
		ctx.Redirect(302,beego.URLFor("FrontoginController.Get"))
	}
}
func CmsLoginFilter(ctx *context.Context)  {
	//判断session是否为空
	cms_user_name :=ctx.Input.Session("cms_user_name")
	if cms_user_name==nil {
		//为空的话就跳转到登录页面
		ctx.Redirect(302,beego.URLFor("FrontoginController.Get"))
	}
}