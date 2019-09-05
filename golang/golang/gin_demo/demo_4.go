package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router :=gin.Default()
	// 访问静态文件路径
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 返回地三方获取的数据
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	// html 渲染
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// 重定向
	router.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	router.GET("/test1", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		router.HandleContext(context)
	})
	router.GET("/test2", func(context *gin.Context) {
		context.JSON(200, gin.H{"hello": "world"})
	})

	// 自定义中间件
	router.Use(Logger())
	router.GET("/test3", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		log.Println(example)
		context.JSON(200, "hello world")
	})
	router.Run("172.18.62.24:3000")
}

func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
