package routes

import (
	"dev-framework-go/conf"
	_ "dev-framework-go/docs"
	diyerror "dev-framework-go/middleware/error"
	"dev-framework-go/middleware/log"
	"dev-framework-go/middleware/session"
	v1 "dev-framework-go/routes/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(log.DiyLogger())
	r.Use(gin.Recovery())
	r.Use(diyerror.CatchError())
	r.Use(session.SessionInit(conf.SESSION_NAME))
	r.GET("/api/v1/login", v1.LoginHandler())
	r.GET("/api/v1/logout", v1.LogOutHandler())
	apiv1 := r.Group("/api/v1")
	apiv1.Use(session.SessionV1())
	{
		apiv1.GET("/record", v1.RecordHandler())
		apiv1.GET("/detail", v1.BillDetailHandler())
		apiv1.GET("/info", v1.UserInfoHandler())
	}
	if conf.OPEN_SWAGGER == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r
}
