package  middleware

import (
	"goblog/utils"
)
import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte(utils.JwtKey)

// 请求结构体
type MyClaims struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	jwt.StandardClaims

}
// 生成token
func SetToken(username string,password string) string{
	//expireTime :=time.Now().Add(10*time.Hour)
	//
	//SetClaims :=MyClaims{
	//
	//}
	return ""
}
// 验证token

// jwt中间件