package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 默认路由
	r := gin.Default()

	// 模板路径
	r.LoadHTMLGlob("templates/*")

	// route配置
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "你好值: %v", "好的")
	})

	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "newa.html", gin.H{
			"title": "我是你爸爸",
		})
	})

	r.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "hello, yes sir",
		})
	})

	r.GET("/json2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "json2",
		})
	})

	// web服务
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
