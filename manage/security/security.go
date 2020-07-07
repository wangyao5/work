package security

import (
	"fmt"
	"manage/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Security 权限控制模块.
type Security struct {
}

//Login 账户登陆.
func (Security) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("passwd")
	userModel := UserModel{}.FindOne(username)
	if userModel.Username == username && userModel.Status == 1 {
		if util.MD5(password+userModel.Salt) == userModel.Password {
			claims := Claims{}
			accessToken, accessTokenErr := CreateAccessToken(claims)
			refreshToken, refreshTokenErr := CreateAccessToken(claims)
			if accessTokenErr == nil && refreshTokenErr == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"accessToken":  accessToken,
					"refreshToken": refreshToken,
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": fmt.Sprintf("%#v\n%#v", accessTokenErr, refreshTokenErr),
			})
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": fmt.Sprintf("username: %s not found", username),
		})
	}
}

//Register 注册商户.
func (Security) Register(ctx *gin.Context) {

}

//ResetPwd 重置商户密码
func (Security) ResetPwd(ctx *gin.Context) {

}
