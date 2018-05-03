package index

import (
	"bee-go-myBlog/controllers"
	"strconv"
	"bee-go-myBlog/services"
	"bee-go-myBlog/models"
	"fmt"
)

type HomeController struct {
	controllers.BaseController
}

//@router / [get]
func (h *HomeController) Index() {
	page := h.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}

	post,err := services.IndexPostList(page2)
	cate := services.IndexAllCateBySort()
	tag := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	totalPage,lastPage,currentPage,nextPage := models.PostPaginate(page2,"index")
	h.Data["totalPage"] = totalPage
	h.Data["lastPage"] = lastPage
	h.Data["currentPage"] = currentPage
	h.Data["nextPage"] = nextPage
	h.Data["post"] = post

	h.Data["system"] = system
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tag


	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
		h.TplName = "home/local/index.tpl"
	} else {
		h.Layout = "home/master.tpl"
		h.TplName = "home/index.tpl"
	}

}

//@router /detail/:id([0-9]+ [get]
func (h *HomeController) Detail() {
	id := h.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	post := services.IndexPostDetail(id64)
	lastPost := services.IndexPostLast(id64)
	lastBefore := services.IndexPostBefore(id64)
	cate := services.IndexAllCateBySort()
	tag := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	if lastPost == nil {
		h.Data["lastPostCond"] = false
	} else {
		h.Data["lastPostCond"] = true
	}
	if lastBefore == nil {
		h.Data["lastBeforeCond"] = false
	} else {
		h.Data["lastBeforeCond"] = true
	}
	postTag := services.IndexPostTag(id64)

	go services.PostReadNumAdd(id64)
	fmt.Println(postTag)
	//评论
	//阅读数+1
	h.Data["system"] = system
	h.Data["cate"] = cate
	h.Data["link"] = link
	h.Data["tag"] = tag
	h.Data["postTag"] = postTag
	h.Data["detail"] = post
	h.Data["lastPost"] = lastPost
	h.Data["lastBefore"] = lastBefore

	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
	} else {
		h.Layout = "home/master.tpl"
	}
	h.TplName = "home/detail.tpl"

}

//@router /categories/:cate [get]
func (h *HomeController) Cate() {
	page := h.GetString("page")
	page2, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		page2 = 1
	}
	cate := h.Ctx.Input.Param(":cate")
	categories,err := services.IndexGetCateByName(cate)
	if err != nil {
		//走404页面
	}
	postCates,err := services.IndexGetCatePost(categories.Id,page2)
	fmt.Println(postCates,err,"234234")
	if err != nil {
	}
	posts,err := services.IndexPostByIds(postCates,page2)
	if err != nil {
	}

	totalPage,lastPage,currentPage,nextPage := models.IndexCatePostPaginate(page2,categories.Id)
	h.Data["totalPage"] = totalPage
	h.Data["lastPage"] = lastPage
	h.Data["currentPage"] = currentPage
	h.Data["nextPage"] = nextPage

	fmt.Println(totalPage,lastPage,currentPage,nextPage,"看结果")
	cates := services.IndexAllCateBySort()
	tag := services.IndexAllTag()
	system := services.IndexSystem()
	link := services.IndexLinkList()
	h.Data["posts"] = posts
	h.Data["system"] = system
	h.Data["cate"] = cates
	h.Data["link"] = link
	h.Data["tag"] = tag

	if system.CdnType == 1 {
		h.Layout = "home/local/master.tpl"
	} else {
		h.Layout = "home/master.tpl"
	}
	h.TplName = "home/cate.tpl"
}

//@router /tags/:id([0-9]+ [get]
func (h *HomeController) Tag() {

}

//@router /archive [get]
func (h *HomeController) Archive() {

}