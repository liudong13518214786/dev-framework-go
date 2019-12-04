package main

import (
	"dev-framework-go/pkg/cache"
	"github.com/gin-gonic/gin"
)

func init() {
	//初始化redis连接池
	cache.InitRedisPool()
	//初始化db连接池

}

func main() {
	r := gin.Default()
	r.GET("")
}
