package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"test_blob/controllers/cms"
   "test_blob/controllers/front"
)

func init() {
   //cms端
   beego.Router("/cms",&cms.LoginController{})
   beego.Router("/cms/main/main",&cms.MainController{})
   beego.Router("/cms/main/post",&cms.PostController{})
   beego.Router("/cms/main/welcome",&cms.MainController{},"get:Welcome")
   beego.Router("/cms/main/post_to_add",&cms.PostController{},"get:ToAdd")
   beego.Router("/cms/main/post_do_add",&cms.PostController{},"post:DoAdd")
   beego.Router("/cms/main/post_delete",&cms.PostController{},"get:PostDelete")
   beego.Router("/cms/main/post_to_edit",&cms.PostController{},"get:ToEdit")
   beego.Router("/cms/main/post_do_edit",&cms.PostController{},"post:DoEdit")
   //beego.Router("/",)
   //前端
   beego.Router("/",&front.IndexController{})
   beego.Router("/detail",&front.IndexController{},"get:PostDetail")
   beego.Router("/register",&front.RegisterController{})
   beego.Router("/login",&front.FrontLoginController{})
   beego.Router("/comment",&front.CommentController{})
}
