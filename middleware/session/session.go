package session

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieManger := util.CookieManger{c, conf.SESSION_NAME, conf.COOKIE_EXPIRE_TIME}
		cookieId := cookieManger.GetSessionid()
		sessionManage := util.SessionManager{cookieId}
		userUuid := sessionManage.Get("useruuid")
		code := conf.SUCCESS
		if userUuid == "" {
			code = conf.NOT_LOGIN
		}
		if code != conf.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  conf.GetMessage(code),
				"data": nil,
			})
			c.Abort()
		}
		c.Next()
	}
}
