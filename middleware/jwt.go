package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/utils"
	"goblog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)
import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte(utils.JwtKey)

// 请求结构体
type MyClaims struct {
	Username string `json:"username"`

	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour).Unix()

	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "xbdh",
		},
	}
    //ES256 出错
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
    //fmt.Println(JwtKey,reqClaim,"---")
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// 验证token

func CheckToken(token string)(*MyClaims,int){
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errmsg.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errmsg.ERROR_TOKEN_RUNTIME
			} else {
				return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, errmsg.SUCCSE
		} else {
			return nil, errmsg.ERROR_TOKEN_WRONG
		}
	}
	return nil, errmsg.ERROR_TOKEN_WRONG
}

// jwt中间件
func JwtToken()gin.HandlerFunc{
	return func(c *gin.Context){
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode != errmsg.SUCCSE {
			code = tCode
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key)
		c.Next()
	}

}