package log

import (
	"dev-framework-go/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

//日志中间件
func DiyLogger() gin.HandlerFunc {
	// 请求日志
	return gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {

		return fmt.Sprintf("[%s] %d %s %s (%s) [%v] %s\n",
			p.TimeStamp.Format(conf.TIME_FORMAT),
			p.StatusCode,
			p.Request.Method,
			p.Path,
			p.ClientIP,
			p.Request.PostForm.Encode(), //参数
			p.Latency,                   //执行时间
		)
	})
}
