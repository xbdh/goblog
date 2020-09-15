package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:archar(20) ;not null" json:"username"`

	Password string `gorm:"type:varchar(20) ;not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUser(name string) (errCode int) {
	var u User
	db.Select("id").Where("username = ?", name).First(&u)
	if u.ID > 0 {
		errCode = errmsg.ERROR_USERNAME_USED
		// 两个return都要加
		return
	}
	errCode = errmsg.SUCCSE
	return
}

// 新增用户
func CreateUser(u *User) (errCode int) {
	err := db.Create(&u).Error
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
func GetUsers(pageSize int, pageNum int) []User {
	var us []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&us).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return us
}


// 编辑用户
func EditUser(){

}