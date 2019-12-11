package v1

import (
	"dev-framework-go/conf"
	"dev-framework-go/models"
	"net/http"
	//"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
)

type LoginParm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//sql := util.SelectSql("project", []string{"uuid", "descs", "build_time"}, nil, "", "", "")
		res := models.GetBillList("10", "0", "")
		//res:=models.SelectTags("select * from project")
		c.JSON(200, gin.H{
			"code": 100,
			"data": res,
		})
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginP LoginParm
		if err := c.ShouldBind(&loginP); err == nil {

		}
		//username:=c.Query("username")
		//password:=c.Query("password")
		//if username == "" || password == ""{
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": conf.MISS_PARAMS,
		//		"msg":conf.GetMessage(conf.MISS_PARAMS),
		//		"data": nil,
		//	})
		//	return
		//}
		c.JSON(http.StatusOK, gin.H{
			"code": conf.SUCCESS,
			"msg":  "",
			"data": loginP,
		})
	}
}
