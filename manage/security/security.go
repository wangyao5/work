package security

import (
	"fmt"
	"manage/util"
	"net/http"
	"time"

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
			//查询API接口
			claims := Claims{UID: username}
			accessTokenExpireTime := time.Minute * 60
			accessToken, accessTokenErr := CreateAccessToken(claims, time.Now().Add(accessTokenExpireTime).Unix())
			refreshTokenExpireTime := time.Hour * 24
			refreshToken, refreshTokenErr := CreateRefreshToken(claims, time.Now().Add(refreshTokenExpireTime).Unix())
			if accessTokenErr == nil && refreshTokenErr == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"access-token":          accessToken,
					"access-token-expired":  accessTokenExpireTime.Seconds(),
					"refresh-token":         refreshToken,
					"refresh-token-expired": refreshTokenExpireTime.Seconds(),
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

//RefreshToken 刷新Token
func (Security) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.PostForm("refresh-token")
	if claims, ok := ValidateRefreshToken(refreshToken); ok {

	}
}
