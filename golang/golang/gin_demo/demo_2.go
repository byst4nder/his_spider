package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"time"
)

func main()  {
	// 无中间件
	//r := gin.New()
	// 默认启动方式，包含Logger、Recovery 中间件
	// r1 :=gin.Default()
	// 全局 中间件
	// 使用Logger 中间件
	//r.Use(gin.Logger())
	// 使用Recovery 中间件
	//r.Use(gin.Recovery())
	// 路由添加中间件， 可以添加任意多个
	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	/*authorized := r.Group("/")
	authorized.Use(AuthRequired()){
		authorized.POST("login", loginEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analuticsEcdppoint)
	}*/

	// 写日志
	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	//讲日志同事写入文件和控制台
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	// 自定义日志格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":3000")
}
