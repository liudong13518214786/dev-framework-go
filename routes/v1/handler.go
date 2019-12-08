package v1

import (
	"github.com/gin-gonic/gin"
)

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Query("name")
		c.JSON(200, gin.H{
			"code": 100,
			"msg":  a,
		})
	}
}
