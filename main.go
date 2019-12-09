package main

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/cache"
	"dev-framework-go/pkg/db"
	"dev-framework-go/routes"
	"github.com/wonderivan/logger"
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
	_ = r.Run(conf.PORT)
}
