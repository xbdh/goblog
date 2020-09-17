package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"goblog/utils/validator"
	"net/http"
	"strconv"
)


var code int

// 添加用户
func AddUser(c *gin.Context) {
	var u model.User
	_ = c.ShouldBindJSON(&u)
	var msg string
	msg, code = validator.Validate(&u)
	if code != errmsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": msg,
			},
		)
		c.Abort()
	}
	code = model.CheckUser(u.Username)
	if code == errmsg.SUCCSE {
		code = model.CreateUser(&u)
	}

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUserInfo(c *gin.Context) {

}

// 查询用户列表 分页
func GetUsers(c *gin.Context) {
	// 不传pagesize和pagenum 查询全部
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	us := model.GetUsers(pageSize, pageNum)
	errCode := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": errmsg.GetErrMsg(errCode),
		"data":    us,
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var u model.User
	_ = c.ShouldBindJSON(&u)
	id, _ := strconv.Atoi(c.Param("id"))
	errCode := model.CheckUser(u.Username)
	if errCode == errmsg.SUCCSE {
		_ = model.EditUser(id, &u)
	}

	if errCode == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": errmsg.GetErrMsg(errCode),
	})
}

// 删除用户

func DeleteUser(c *gin.Context) {
	// param 和query 一定要分清啊
	id, _ := strconv.Atoi(c.Param("id"))
	errCode := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": errmsg.GetErrMsg(errCode),
	})

}
