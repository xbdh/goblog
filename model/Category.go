package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm :"type: varchar(20);not null" json:"name"`

}

func CheckCategory (name string) (errCode int) {
	var c Category
	db.Select("id").Where("name = ?", name).First(&c)
	if c.ID > 0 {
		errCode = errmsg.ERROR_CATENAME_USED
		// 两个return都要加
		return
	}
	errCode = errmsg.SUCCSE
	return
}

// 新增用户
func CreateCategory (c *Category ) (errCode int) {


	err := db.Create(&c).Error

	if err != nil {
		errCode = errmsg.ERROR
		//return errmsg.ERROR
		return
	}

	errCode = errmsg.SUCCSE
	return
	//return errmsg.SUCCSE
}

// 查询用户列表
func GetCategorys(pageSize int, pageNum int) []Category{
	var cs []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cs
}


// 编辑用户
//
func EditCategory (id int,c * Category) (errCode int){
	var maps = make(map[string]interface{})
	maps["name"]=c.Name

	err = db.Model(&Category{}).Where("id =? ",id).Update(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}

	return errmsg.SUCCSE

}
// 删除用户
func DeleteCategory (id int)(errCode int){
	var c Category
	// gorm 操作待看
	err =db.Where("id= ? ",id).Delete(&c).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}