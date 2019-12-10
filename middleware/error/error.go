package error

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

func CatchError() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("[PINIC ERROR]", r)
				//context.Abort()
				context.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "系统异常",
					"data": nil,
				})
			}
		}()
		context.Next()
	}
}
