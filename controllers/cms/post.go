package cms

import (

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"test_blob/models"
	"test_blob/utils"
	"fmt"
	"time"
	"strconv"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Get()  {

	o := orm.NewOrm()
	posts := []models.Post{}
	qs := o.QueryTable(new(models.Post))

   //1对多需要关联查询
	qs.RelatedSel().All(&posts)
	//分页操作
	//count为数据库的条数
	count,_ := qs.Count()
	//获取当前页面 默认的或者是读者点击的
	current_page,err := p.GetInt("p")
	if err != nil {
		current_page = 1
	}
	//页数
	page_size := 10
	//每一页的条数
	total_pages := utils.GetPageNum(count,page_size)
	arround_count := 4
	//Get_pagination_data 函数里的输入含义为 总页数，当前页，前后页
	//判断left_has_more；left_has_more 左右边的页码，有就显示
	left_pages, right_pages, left_has_more, right_has_more := utils.Get_pagination_data(total_pages,current_page,arround_count)

	has_pre_page, has_next_page, pre_page, next_page := utils.HasNext(current_page, total_pages)

	p.Data["left_pages"] = left_pages
	p.Data["left_has_more"] = left_has_more
	p.Data["page"] = current_page

	p.Data["has_pre_page"] = has_pre_page
	p.Data["pre_page"] = pre_page
	p.Data["has_next_page"] = has_next_page
	p.Data["next_page"] = next_page

	p.Data["right_pages"] = right_pages

	p.Data["right_has_more"] = right_has_more


	p.Data["num_pages"] = total_pages //总页数
	p.Data["count"] = count //总数量
	p.Data["posts"] = posts
	p.TplName = "cms/post-list.html"

}

func (p *PostController) ToAdd()  {
	p.TplName = "cms/post-add.html"
	
}
//获取添加帖子
func (p *PostController) DoAdd()  {
//获取后端返回的数据
	title := p.GetString("title")
	desc := p.GetString("desc")
	content := p.GetString("content")
	f,h,err := p.GetFile("cover")

	defer f.Close()

	var cover string
	if err != nil {
		cover = "static/upload/no_pic.jpg"
	}

	// 生成时间戳，防止重名
	timeUnix:=time.Now().Unix() // int64类型
	time_str := strconv.FormatInt(timeUnix,10) // 将int64转为字符串 convert：转换

	path := "static/upload/"+time_str+h.Filename
	// 保存获取到的文件
	err1 := p.SaveToFile("cover",path)

	if err1 != nil {
		cover = "static/upload/no_pic.jpg"
	}
	cover = path

	o := orm.NewOrm()

	author := p.GetSession("cms_user_name")
	user := models.User{}
	o.QueryTable(new(models.User)).Filter("user_name",author).One(&user)
	post := models.Post{
		Title:title,
		Desc:desc,
		Content:content,
		Cover:cover,
		Author:&user,
	}
	//添加到数据库
	_, err2 := o.Insert(&post)

	if err2 != nil {
		fmt.Println("=============")
		fmt.Println(err2)
		p.Data["json"] = map[string]interface{}{"code":500,"msg":err2}
		p.ServeJSON()
	}


	p.Data["json"] = map[string]interface{}{"code":200,"msg":"添加成功"}
	p.ServeJSON()
	
}
//删除帖子
func (p *PostController) PostDelete()  {

	id,err := p.GetInt("id")
	if err != nil {
		p.Ctx.WriteString("id参数错误")
	}
	o := orm.NewOrm()
	_,err2 := o.QueryTable(new(models.Post)).Filter("id",id).Delete()

	if err2 != nil {
		fmt.Println(err2)
		p.Ctx.WriteString("删除错误")
	}
//跳转页面
	p.Redirect(beego.URLFor("PostController.Get"),302)

}

func (p *PostController) ToEdit()  {

	id,err := p.GetInt("id")
	if err != nil {
		p.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()

	post := models.Post{}
	o.QueryTable(new(models.Post)).Filter("id",id).One(&post)
	p.Data["post"] = post
	p.TplName = "cms/post-edit.html"

}

func (p *PostController) DoEdit()  {

	o := orm.NewOrm()


	id,err := p.GetInt("id")
	if err !=nil {
		p.Data["json"] = map[string]interface{}{"code":500,"msg":"id参数错误"}
	}

	qs := o.QueryTable(new(models.Post)).Filter("id",id)


	title := p.GetString("title")
	desc := p.GetString("desc")
	content := p.GetString("content")



	//获取文件
	f,h,err1 := p.GetFile("cover")


	fmt.Println("==============")
	fmt.Println(id)
	fmt.Println(title)
	fmt.Println(desc)
	fmt.Println(content)
	fmt.Println(err1)


	if err1 != nil {

		_,err4 := qs.Update(orm.Params{
			"title":title,
			"desc":desc,
			"content":content,
		})

		if err4 != nil {
			p.Data["json"] = map[string]interface{}{"code":500,"msg":"更新失败"}
		}
		p.Data["json"] = map[string]interface{}{"code":200,"msg":"更新成功"}
		p.ServeJSON()


	}


	defer f.Close()


	// 生成时间戳，防止重名
	timeUnix:=time.Now().Unix() // int64类型
	time_str := strconv.FormatInt(timeUnix,10) // 将int64转为字符串 convert：转换

	path := "static/upload/"+time_str+h.Filename
	// 保存获取到的文件
	err2 := p.SaveToFile("cover",path)

	if err2 != nil {
		_,err5 := qs.Update(orm.Params{
			"title":title,
			"desc":desc,
			"content":content,
		})
		if err5 != nil {
			p.Data["json"] = map[string]interface{}{"code":500,"msg":"更新失败"}
		}
		p.Data["json"] = map[string]interface{}{"code":200,"msg":"更新成功"}
		p.ServeJSON()
	}
	_,err6 := qs.Update(orm.Params{
		"title":title,
		"desc":desc,
		"content":content,
		"cover":path,
	})

	if err6 != nil {
		p.Data["json"] = map[string]interface{}{"code":500,"msg":"更新失败"}
	}
	p.Data["json"] = map[string]interface{}{"code":200,"msg":"更新成功"}
	p.ServeJSON()



}

