package services

import (
	"bee-go-myBlog/models"

	"bee-go-myBlog/helper"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func GetMyAllCate() {

}

func GetAllCateBySort() ([]interface{}) {
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64

	fields = []string{"Id", "Name", "DisplayName", "ParentId", "Description","CreatedAt"}
	cate, _ := models.GetAllCategories(query, fields, sortby, order, offset, limit)
	if len(cate) == 0 {
		cateCreate := &models.Categories{
			Name		:	"default",
			DisplayName	:	"默认分类",
			Description	:	"默认生成的顶级分类",
			ParentId	:	0,
		}
		models.AddCategories(cateCreate)
		cate, _ = models.GetAllCategories(query, fields, sortby, order, offset, limit)
	}
	return tree(cate, 0, 0,0)
}

func GetCateByLike(param string) (ml map[int64]string) {
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64

	query = make(map[string]string)
	if param != "" {
		query["display_name__icontains"] = param
	}

	fields = []string{"Id", "Name", "DisplayName", "ParentId", "Description"}
	cates,_ := models.GetAllCategories(query, fields, sortby, order, offset, limit)
	cate := make(map[int64]string, len(cates))
	for _, v := range cates{
		cate[v.(map[string]interface{})["Id"].(int64)] = v.(map[string]interface{})["DisplayName"].(string)
	}

	return cate
}


func GetCateByPostId(postId int64) (int64,error) {
	o := orm.NewOrm()
	artCate := models.ArticleCate{ArtId:postId}
	err := o.Read(&artCate,"ArtId")
	if err == orm.ErrNoRows {
		return 0,err
	}
	return artCate.CateId,nil
}

func tree(cate []interface{}, parent int64, level int64,key int64) ([]interface{}) {
	html := "-"
	var data []interface{}
	for _, v := range cate {
		var ParentId = v.(map[string]interface{})["ParentId"].(int64)
		var Id = v.(map[string]interface{})["Id"].(int64)
		if ParentId == parent {
			var newHtml string
			if level != 0 {
				newHtml = helper.GoRepeat("&nbsp;&nbsp;&nbsp;&nbsp;", level) + "|"
			}
			v.(map[string]interface{})["html"] = newHtml + helper.GoRepeat(html, level)
			data = append(data,v)
			data = helper.GoMerge(data,tree(cate, Id, level+1,key+1))
		}
	}
	return data
}

func GetSimilar(beginId int64) {
	o := orm.NewOrm()
	cates := models.Categories{ParentId:beginId}
	err := o.Read(&cates,"ParentId")
	if err != nil {
		return
	}
	fmt.Println(2343)
	//GetChild(cates)

}

func GetChild(cates []interface{}) {
	//for _,v := range cates {
	//
	//}
}