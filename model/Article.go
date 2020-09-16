package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type Article struct {
	Category Category`gorm:"foreignKey:Cid"`
	gorm.Model

	Title string`gorm:"type:varchar(100);not null" json:"title"`
	Cid int`gorm:"type :int ;not null" json:"cid"`
	Desc string`gorm:"type :varchar(200)" json:"desc"`
	Content string`gorm:"type:longtext" json:"content"`
	Img string`gorm:"type:varchar(100)" json:"img"`
}



// 新增文章
func CreateArticle (art *Article) (errCode int) {


	err := db.Create(&art).Error

	if err != nil {
		errCode = errmsg.ERROR
		//return errmsg.ERROR
		return
	}

	errCode = errmsg.SUCCSE
	return
	//return errmsg.SUCCSE
}

// 某分类下的所有文章
func GetCateOfArticle(id int,pageSize int, pageNum int) ([]Article,int){
	var arts []Article
	errCode :=db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Where("cid = ?",id).Find(&arts).Error

	if errCode!=nil{
		return nil,errmsg.ERROR_CATE_NOT_EXIST
	}
	return arts,errmsg.SUCCSE
}

// 查询单个文章
func GetArticleInfo(id int)  (Article,int) {
	var art Article
	err:=db.Preload("Category").Where("id = ?",id).First(&art).Error
	if err !=nil{
		return art,errmsg.ERROR_ART_NOT_EXIST
	}
	return art ,errmsg.SUCCSE
}
// 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article,int){
	var arts []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,errmsg.ERROR
	}
	return arts,errmsg.SUCCSE
}


// 编辑文章
//
func EditArticle (id int,art *Article) (errCode int){
	var maps = make(map[string]interface{})
	maps["title"]=art.Title
	maps["cid"]=art.Cid
	maps["desc"]=art.Desc
	maps["content"]=art.Content
	maps["img"]=art.Img



	err = db.Model(&Article{}).Where("id =? ",id).Update(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}

	return errmsg.SUCCSE

}
// 删除文章
func DeleteArticle (id int)(errCode int){
	var art Article
	// gorm 操作待看
	err =db.Where("id= ? ",id).Delete(&art).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}