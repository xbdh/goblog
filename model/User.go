package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20) ;not null" json:"username"`

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

	u.Password =ScryptPassword(u.Password)
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
// 更新除密码以外的字段，字段类型不同
func EditUser(id int,u *User  ) (errCode int){
	var maps = make(map[string]interface{})
	maps["username"]=u.Username
	maps["role"]=u.Role
	err = db.Model(&User{}).Where("id =? ",id).Update(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}

	return errmsg.SUCCSE

}
// 删除用户
func DeleteUser(id int)(errCode int){
	var u User
	// gorm 操作待看
	err =db.Where("id= ? ",id).Delete(&u).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 密码加密
// bcrypt scrypt 两种加密，也可以用钩子函数

func ScryptPassword(password string)string  {
	//返回的密码【】byte 的长度
	const keyLen=15
	salt:=make([]byte,8)

	salt =[]byte{20,20,9,16,12,46}
	HashPassword,err:=scrypt.Key([]byte(password),salt,32768, 8, 1, keyLen)
	if err!=nil{
		log.Fatal(err)
	}
	//fmt.Println(HashPassword)

	fpw:=base64.StdEncoding.EncodeToString(HashPassword)
	return fpw
}

// 钩子函数
//func (u* User)BeforeSave(){
//
//}