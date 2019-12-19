package routes

import (
	diyerror "dev-framework-go/middleware/error"
	"dev-framework-go/middleware/log"
	v1 "dev-framework-go/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(log.DiyLogger())
	r.Use(gin.Recovery())
	r.Use(diyerror.CatchError())
	r.GET("/api/v1/login", v1.LoginHandler())
	r.GET("/api/v1/logout", v1.LogOutHandler())
	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/record", v1.RecordHandler())
		apiv1.GET("/detail", v1.BillDetailHandler())
	}
	return r
}
