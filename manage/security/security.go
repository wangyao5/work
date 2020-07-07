package security

import "github.com/gin-gonic/gin"

//Security 权限控制模块.
type Security struct {
}

//Login 账户登陆.
func (Security) Login(ctx *gin.Context) {

}

//Register 注册商户.
func (Security) Register(ctx *gin.Context) {

}

//ResetPwd 重置商户密码
func (Security) ResetPwd(ctx *gin.Context) {

}
