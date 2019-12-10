package error

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"runtime/debug"
	"strings"
)

func CatchError() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("[PINIC ERROR]", r)
				//context.Abort()
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				requestUrl := fmt.Sprintf("%s  %s%s", context.Request.Method, context.Request.Host, context.Request.RequestURI)
				emailInfo := conf.EmailDetail
				body := strings.ReplaceAll(emailInfo, "[error_info]", fmt.Sprintf("%s", r))
				body = strings.ReplaceAll(body, "[request_time]", util.GetNowTime())
				body = strings.ReplaceAll(body, "[request_url]", requestUrl)
				body = strings.ReplaceAll(body, "[request_ua]", context.Request.UserAgent())
				body = strings.ReplaceAll(body, "[error_debug]", DebugStack)

				_ = util.SendMail(fmt.Sprintf("【重要错误】%s 项目出错了！", conf.APPNAME), body)
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
