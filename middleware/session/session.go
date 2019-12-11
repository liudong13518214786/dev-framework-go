package session

import (
	"dev-framework-go/conf"
	s "dev-framework-go/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

//验证登录中间件
func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		userUuid := s.SessionGet(c, "useruuid")
		//cookieManger := util.CookieManger{c, conf.SESSION_NAME, conf.COOKIE_EXPIRE_TIME}
		//cookieId := cookieManger.GetSessionid()
		//sessionManage := util.SessionManager{cookieId}
		//userUuid := session.Get("useruuid")
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
