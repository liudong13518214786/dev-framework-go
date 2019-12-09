package v1

import (
	"dev-framework-go/models"
	"github.com/gin-gonic/gin"
)

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Query("name")
		res := models.SelectTags("select uuid,pj_name,show_name,descs,build_time, status from project")
		//res:=models.SelectTags("select * from project")
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  a,
			"data": res,
		})
	}
}
