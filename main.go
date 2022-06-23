package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "Page Not Found")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"code": "404", "message": "Page Not Found"})
		}
	})

	// 运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")
}
