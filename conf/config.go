package conf

import "github.com/gin-gonic/gin"

const (
	APPNAME                 = "开发框架测试"
	JWTTOKEN                = "123123"
	REDIS_NETWORK           = "tcp"
	REDIS_ADDRESS           = "dbi.mshare.cn:1200"
	REDIS_PASS              = "255ZDVsni98AB4KBHd76pyAh"
	REDIS_MAXIDLE           = 10
	REDIS_MAXACTIVE         = 10
	REDIS_IDLETIMEOUT       = 10
	REDIS_WAIT         bool = true
	SESSION_NAME            = "SXS-TEST"
	COOKIE_SECRET           = "dfafafa"
	DOMAIN                  = "localhost"
	COOKIE_EXPIRE_TIME      = 600
	PORT                    = ":8890"
	APP_ENV                 = "release"
	TIME_FORMAT             = "2006-01-02 15:04:05"
	DB_HOST                 = "dbi.mshare.cn"
	DB_PORT                 = 1094
	DB_USER                 = "dbuser"
	DB_PASS                 = "dY8*6fN6Z#xSOg$wG9zDATTe"
	DB_NAME                 = "mx_bill"
	DB_MaxOpenConns         = 100
	DB_MaxIdleConns         = 10 //空闲时的最大连接数
	ERRORNOTIFYOPEN    bool = false
	SYSTEMEMAILUSER         = "liudong@mshare.cn"
	EMAILTOUSER             = "846723063@qq.com" //多个用户逗号隔开
	SYSTEMEMAILHOST         = "smtp.mxhichina.com"
	SYSTEMEMAILPORT         = 25
	SYSTEMEMAILPASS         = "LD@qq.com"
	PERNUM                  = 10 //每页显示条数
)

func InitAppEnv(env string) {
	switch env {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
