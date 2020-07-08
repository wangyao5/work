package filter

import (
	"manage/security"
	"strings"

	"github.com/gin-gonic/gin"
)

//Authorization 身份认证中间件.
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := ctx.GetHeader("authorization")
		if claims, ok := security.ValidateAccessToken(accessToken); ok {
			// claims.API
			paths := strings.Split(claims.API, ",")
			for _, path := range paths {
				if strings.HasPrefix(c.Request.RequestURI, path) {
					c.Next()
					return
				}
			}
		}

		// 处理请求
		c.JSON(http.Statusok, gin.H{
			"msg": "authorization failed",
		})
	}
}
