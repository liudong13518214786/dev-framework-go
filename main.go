package main

import (
	"dev-framework-go/pkg/cache"
	"dev-framework-go/routes"
)

func init() {
	//初始化redis连接池
	cache.InitRedisPool()
	//初始化db连接池

}

func main() {
	r := routes.InitRoute()
	_ = r.Run()
}
