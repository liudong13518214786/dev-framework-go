package routes

import (
	"dev-framework-go/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tag")
	}
	return r
}
