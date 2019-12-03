package jwt

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
