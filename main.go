package main

import "github.com/gin-gonic/gin"

func main() {

	// 默认路由
	r := gin.Default()

	// route配置
	r.GET("/", func(context *gin.Context) {
		context.String(200, "你好值: %v", "好的")
	})

	// web服务
	r.Run()
}
