package main

import (
	"dev-framework-go/pkg/cache"
	"github.com/gin-gonic/gin"
)

func init() {
	cache.InitRedisPool()
}

func main() {
	r := gin.Default()
	r.GET("")
}
