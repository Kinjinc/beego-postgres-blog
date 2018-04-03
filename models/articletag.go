package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ArticleTag struct {
	Id     	int64 `orm:"column(id);auto;unique" json:"id"`
	ArtId 	int64 `orm:"column(art_id);" json:"art_id"`
	TagId 	int64 `orm:"column(tag_id);" json:"tag_id"`
}

func init() {
	orm.RegisterModel(new(ArticleTag))
}

func (a *ArticleTag) TableName() string {
	return "y_article_tag"
}

// AddArticleTag insert a new ArticleTag into database and returns
// last inserted Id on success.
func AddArticleTag(m *ArticleTag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticleTagById retrieves ArticleTag by Id. Returns error if
// Id doesn't exist
func GetArticleTagById(id int64) (v *ArticleTag, err error) {
	o := orm.NewOrm()
	v = &ArticleTag{Id: id}
	if err = o.QueryTable(new(ArticleTag)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticleTag retrieves all ArticleTag matches certain condition. Returns empty list if
// no records exist
func GetAllArticleTag(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ArticleTag))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ArticleTag
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateArticleTag updates ArticleTag by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleTagById(m *ArticleTag) (err error) {
	o := orm.NewOrm()
	v := ArticleTag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticleTag deletes ArticleTag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticleTag(id int64) (err error) {
	o := orm.NewOrm()
	v := ArticleTag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ArticleTag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func GetTagIdByPostId(postId int64) (num int64,maps []orm.Params)  {
	o := orm.NewOrm()
	num,_ = o.QueryTable(new(ArticleTag)).Filter("ArtId",postId).Values(&maps)
	return num,maps
}