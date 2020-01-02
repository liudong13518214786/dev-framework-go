package routes

import (
	"dev-framework-go/conf"
	_ "dev-framework-go/docs"
	"dev-framework-go/middleware/cross"
	diyerror "dev-framework-go/middleware/error"
	"dev-framework-go/middleware/log"
	"dev-framework-go/middleware/session"
	v1 "dev-framework-go/routes/v1"
	v2 "dev-framework-go/routes/v2"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(log.DiyLogger())
	r.Use(gin.Recovery())
	r.Use(diyerror.CatchError())
	r.Use(session.SessionInit(conf.SESSION_NAME))
	r.Use(cross.Cors())
	//r.LoadHTMLGlob("dist/*")
	r.Static("/static", "./static")
	r.GET("/api/v1/login", v1.LoginHandler())
	r.GET("/api/v1/logout", v1.LogOutHandler())
	r.POST("/api/v1/upload", v2.UploadHandler())

	r.GET("/api/v1/blog", v2.GetBlogListHandler())
	r.GET("/api/v1/detail", v2.DetailHandler())
	r.GET("/api/v1/cate", v2.GetClassHandler())
	apiv1 := r.Group("/api/v1")
	apiv1.Use(session.SessionV1())
	{
		apiv1.GET("/user", v1.UserInfoHandler())
		apiv1.POST("/write", v2.WriteBlogHandler())
	}
	r.GET("/", v2.IndexHandler())
	if os.Getenv("APP_ENV") != "release" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r
}
