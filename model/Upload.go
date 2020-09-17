package model
//
//import (
//	"github.com/aliyun/aliyun-oss-go-sdk/oss"
//	"goblog/utils"
//	"io"
//	"log"
//	"mime/multipart"
//)
//
//func UpLoadFile(file*multipart.File,fileheader *multipart.FileHeader){
//	//// Endpoint以杭州为例，其它Region请按实际情况填写。
//	//endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
//	//// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。
//	//accessKeyId := "<yourAccessKeyId>"
//	//accessKeySecret := "<yourAccessKeySecret>"
//	//bucketName := "<yourBucketName>"
//	//// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
//	objectName := "test/"
//	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
//	localFileName := "<yourLocalFileName>"
//	// 创建OSSClient实例。
//	client, err := oss.New(utils.EndPoint, utils.AccessKeyID, utils.AccessKeySecret)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 获取存储空间。
//	bucket, err := client.Bucket(utils.Bucket)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 上传文件。
//	err = bucket.FPutObject(objectName,multipart)
//	if err != nil {
//		log.Fatal(err)
//	}
//}