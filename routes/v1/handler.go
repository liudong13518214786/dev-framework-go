package v1

import (
	"dev-framework-go/conf"
	s "dev-framework-go/pkg/session"
	"net/http"
	//"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
)

//
//type LoginParm struct {
//	Username string `form:"username"`
//	Password string `form:"password"`
//}

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		Useruuid := s.SessionGet(c, "useruuid")
		c.JSON(200, gin.H{
			"code": 100,
			"data": Useruuid,
		})
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if username == "" || password == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": conf.MISS_PARAMS,
				"msg":  conf.GetMessage(conf.MISS_PARAMS),
				"data": nil,
			})
			return
		}
		Useruuid := s.SessionGet(c, "useruuid")
		if Useruuid == "" {
			s.SessionSet(c, "useruuid", username, conf.COOKIE_EXPIRE_TIME)
		}
		//todo 一边设置session一边取session,第一次会取不到
		Useruuid = s.SessionGet(c, "useruuid")
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  Useruuid,
			"data": nil,
		})
	}
}
