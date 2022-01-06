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

	r.GET("/get", func(context *gin.Context) {
		username := context.Query("username")
		age := context.DefaultQuery("age", "1")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
		})
	})

	r.POST("/post", func(context *gin.Context) {

		username := context.PostForm("username")
		// 页面需要一个form
		context.HTML(http.StatusOK, "newa.html", gin.H{
			"title": username,
		})

		// 获取request body的数据
		// data, _ := context.GetRawData()
	})

	// web服务
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
