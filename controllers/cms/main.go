package cms

import beego "github.com/beego/beego/v2/server/web"

type MainController struct {
	beego.Controller
}

func (m *MainController) Get()  {
	m.TplName="cms/index.html"
}
func (m *MainController) Welcome() {
	m.TplName="cms/welcome.html"
}