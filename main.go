package main

import (
	"context"
	"dev-framework-go/conf"
	"dev-framework-go/pkg/cache"
	"dev-framework-go/pkg/db"
	"dev-framework-go/routes"
	"github.com/wonderivan/logger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	// todo 初始化用户级别的日志
	//初始化环境
	conf.InitAppEnv(conf.APP_ENV)
	//初始化redis连接池
	cache.InitRedisPool()
	//初始化db连接池
	db.InitDatabasePool()
}

func main() {
	r := routes.InitRoute()
	logger.Debug("SERVER RUN IN http://127.0.0.1" + conf.PORT)
	srv := http.Server{
		Addr:    conf.PORT,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Debug("Server Shotdown...")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Debug("Server exiting")
}
