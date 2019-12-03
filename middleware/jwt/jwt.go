package jwt

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
iss：发行人
exp：到期时间
sub：主题
aud：用户
nbf：在此之前不可用
iat：发布时间
jti：JWT ID用于标识该JWT
除以上默认字段外，我们还可以自定义私有字段，如下例：
{
"sub": "1234567890",
"name": "chongchong",
"admin": true
}
*/
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		code := conf.SUCCESS
		if token == "" {
			code = conf.INVALID_PARAMS
		} else {
			res := util.VerifyToken(token)
			if res {
				code = conf.SUCCESS
			} else {
				code = conf.ValidationError
			}
		}
		if code != conf.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  conf.GetMessage(code),
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
