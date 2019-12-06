package routes

import (
	v1 "dev-framework-go/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/tag", v1.TestHandler())
	}
	return r
}
