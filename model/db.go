package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goblog/utils"
	"log"
	"time"
)

var db *gorm.DB
var err error

func InitDb(){
	db,err =gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err !=nil{
		fmt.Printf("连接数据库出错")
		log.Fatal(err)
	}
	// 禁用默认表名的复数形式
	db.SingularTable(true)

	db.AutoMigrate(&User{},&Category{},&Article{})
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//不要加defer 待看
	//defer db.Close()
}