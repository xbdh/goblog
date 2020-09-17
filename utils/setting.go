package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)
var (
	AppMode string
	HttpPort string
	JwtKey  string


	Db string
	DbHost string
	DbPort string
	DbUser string
    DbPassWord string
	DbName string

	EndPoint string
	Bucket string
	AccessKeyID string
	AccessKeySecret string

)

func init(){
	file,err := ini.Load("config/config.ini")
	if err !=nil{
		fmt.Println("路径错误",err)
	}
	LoadServe(file)
	LoadData(file)
	LoadOss(file)
}

func LoadServe(file * ini.File){
	AppMode =file.Section("server").Key("AppMode").MustString("debug")
	HttpPort =file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey =file.Section("server").Key("JwtKey").MustString("wsszzgwdslj")

}
func LoadData(file * ini.File){
	Db =file.Section("database").Key("Db").MustString("mysql")
	DbHost =file.Section("database").Key("DbHost").MustString("localhost")
	DbPort =file.Section("database").Key("DbPort").MustString("3306")
	DbUser =file.Section("database").Key("DbUser").MustString("chen")
	DbPassWord=file.Section("database").Key("DbPassWord").MustString("")
	DbName =file.Section("database").Key("DbName").MustString("goblog")

}

func LoadOss(file * ini.File){
	EndPoint =file.Section("oss").Key("EndPoint").MustString("http://oss-cn-beijing.aliyuncs.com")
	Bucket =file.Section("oss").Key("Bucket").MustString("ginblog")
	AccessKeyID =file.Section("oss").Key("AccessKeyID").MustString("LTAI4FjLbaRPyz3AVBTYmEzv")
	AccessKeySecret =file.Section("oss").Key("AccessKeySecret").MustString("BW2twkD42tPSaOHYxKwBStIRdt3uy")
}