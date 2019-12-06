package log

import "github.com/gin-gonic/gin"

//日志中间件
func DiyLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return ""
	})
}
