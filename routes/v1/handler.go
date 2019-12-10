package v1

import (
	"dev-framework-go/models"
	"fmt"

	//"dev-framework-go/pkg/util"
	"github.com/gin-gonic/gin"
)

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := []string{"123"}

		fmt.Println(a[:len(a)-3])
		//sql := util.SelectSql("project", []string{"uuid", "descs", "build_time"}, nil, "", "", "")
		res := models.GetBillList("10", "0", "")
		//res:=models.SelectTags("select * from project")
		c.JSON(200, gin.H{
			"code": 100,
			"data": res,
		})
	}
}
