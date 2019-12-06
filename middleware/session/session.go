package session

import (
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Cookie()
	}
}
