package security

import (
	"manage/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Security 权限控制模块.
type Security struct {
	userModel *UserModel
}

//Login 账户登陆.
func (s *Security) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("passwd")
	userModel := s.userModel.FindOne()
	if userModel.Username == username && userModel.Status == 1 {
		if util.MD5(password+userModel.Salt) == userModel.Password {

		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

//Register 注册商户.
func (Security) Register(ctx *gin.Context) {

}

//ResetPwd 重置商户密码
func (Security) ResetPwd(ctx *gin.Context) {

}
