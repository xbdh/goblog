package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

func CheckUser(c *gin.Context)  {

}
// 添加用户
func AddUser(c*gin.Context)  {
	var u model.User
	_ =c.ShouldBindJSON(&u)

	errCode :=model.CheckUser(u.Username)
	if errCode==errmsg.SUCCSE{
		_ =model.CreateUser(&u)
	}

	if errCode==errmsg.ERROR_USERNAME_USED{
		errCode =errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":u,
	})
}

// 查询单个用户
func GetUserInfo(c*gin.Context)  {

}

// 查询用户列表 分页
func GetUsers(c *gin.Context){
	// 不传pagesize和pagenum 查询全部
    pageSize ,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum ,_:=strconv.Atoi(c.Query("pagenum"))

	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=-1
	}

	us:= model.GetUsers(pageSize,pageNum)
    errCode:=errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":us,
	})
}
// 编辑用户
func EditUser(c *gin.Context){

}
// 删除用户

func DeleteUser(c *gin.Context)  {

}