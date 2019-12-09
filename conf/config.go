package conf

import "github.com/gin-gonic/gin"

const (
	JWTTOKEN                = "123123"
	REDIS_NETWORK           = "tcp"
	REDIS_ADDRESS           = "127.0.0.1:6379"
	REDIS_MAXIDLE           = 10
	REDIS_MAXACTIVE         = 10
	REDIS_IDLETIMEOUT       = 10
	REDIS_WAIT         bool = true
	SESSION_NAME            = "SXS-TEST"
	COOKIE_SECRET           = "dfafafa"
	DOMAIN                  = "localhost"
	COOKIE_EXPIRE_TIME      = 60
	PORT                    = ":8890"
	APP_ENV                 = "release"
	TIME_FORMAT             = "2006-01-02 15:04:05"
	DB_HOST                 = "dbi.mshare.cn"
	DB_PORT                 = 1094
	DB_USER                 = "dbuser"
	DB_PASS                 = "dY8*6fN6Z#xSOg$wG9zDATTe"
	DB_NAME                 = "sxsdb"
	DB_MaxOpenConns         = 10
	DB_MaxIdleConns         = 10
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
