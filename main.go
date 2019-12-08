package main

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/cache"
	"dev-framework-go/routes"
)

func init() {
	//初始化环境
	conf.InitAppEnv(conf.APP_ENV)
	//初始化redis连接池
	cache.InitRedisPool()
	//初始化db连接池

}

func main() {
	r := routes.InitRoute()
	_ = r.Run(conf.PORT)
}
