package conf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	APPNAME                 = "开发框架测试"
	JWTTOKEN                = "123123"
	REDIS_NETWORK           = "tcp"
	REDIS_ADDRESS           = "redis:6379"
	REDIS_PASS              = "123456"
	REDIS_MAXIDLE           = 10
	REDIS_MAXACTIVE         = 10
	REDIS_IDLETIMEOUT       = 10
	REDIS_WAIT         bool = true
	SESSION_NAME            = "SXS-TEST"
	COOKIE_SECRET           = "PxzKZsak1JFBMg7of0tOCUrG9QYRiv3X"
	COOKIE_EXPIRE_TIME      = 60
	PORT                    = ":8890"
	TIME_FORMAT             = "2006-01-02 15:04:05"
	//DB_HOST                 = "dbi.mshare.cn"
	DB_HOST = "db"
	//DB_PORT                 = 1094
	DB_PORT = 5432
	//DB_USER                 = "dbuser"
	DB_USER = "postgres"
	//DB_PASS                 = "dY8*6fN6Z#xSOg$wG9zDATTe"
	DB_PASS = "123456"
	//DB_NAME                 = "mx_bill"
	DB_NAME              = "users"
	DB_MaxOpenConns      = 100
	DB_MaxIdleConns      = 10 //空闲时的最大连接数
	ERRORNOTIFYOPEN bool = false
	SYSTEMEMAILUSER      = "liudong@mshare.cn"
	EMAILTOUSER          = "846723063@qq.com" //多个用户逗号隔开
	SYSTEMEMAILHOST      = "smtp.mxhichina.com"
	SYSTEMEMAILPORT      = 25
	SYSTEMEMAILPASS      = "LD@qq.com"
	PERNUM               = 5 //每页显示条数
	GLOBAL_SESSION       = "global_session"
	PG_SQL_PRINT         = true
	UploadDst            = "/go/src/dev-framework-go/static/img/"
	PicUrlHost           = "http://127.0.0.1:8890/static/img/"
)

func InitAppEnv() {
	env := os.Getenv("APP_ENV")
	fmt.Printf("this is %s\n", env)
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
