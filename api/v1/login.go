package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context){
	var u  model.User
	c.ShouldBindJSON(&u)

	var token string
	var code int
	code = model.CheckLogin(u.Username, u.Password)
	// fmt.Println(code,"密码")
	if code == errmsg.SUCCSE{
		token ,code =middleware.SetToken(u.Username)
		// fmt.Println(code,"token",token,"token")
	}

	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"message":errmsg.GetErrMsg(code),
		"token" :token,
	})
}