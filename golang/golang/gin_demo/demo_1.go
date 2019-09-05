package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
	"net/http"
)


func main(){
	router := gin.Default()
	// 返回json数据
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello",
		})
	})
	/* 获取路径中的参数
	*/
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello %s", name)
	})
	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(context *gin.Context) {
			name := context.Param("name")
		action := context.Param("action")
		message := name + "  is  " + action
		context.String(http.StatusOK, message)
	})
	// 获取GET参数
	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(context *gin.Context) {
		firstname :=context.DefaultQuery("firstname","Guest")
		lastname := context.Query("lastname") // c.Request.URL.Query().Get("lastname") 的简写
		context.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	// 获取post参数
	router.POST("/form_post", func(context *gin.Context) {
		message :=context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous") // 设置默认值
		context.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	// Get + Post 混合
	// POST /post?id=1234&page=1 HTTP/1.1
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page :=context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")
		context.JSON(http.StatusOK, gin.H{
			"id": id,
			"page": page,
			"name": name,
			"message": message,
		})
	})

	// 上传文件
	router.POST("/upload", func(context *gin.Context) {
		// 单文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		// 上传文件到指定的路径
		dst := "/c/ada"
		context.SaveUploadedFile(file, dst)
		context.String(http.StatusOK, fmt.Printf("%s uploaded", file.Filename ))
	})
	// 多个文件共同上传
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	// 路由分组
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
		v1.POST("/submit")
		v1.POST("/read")
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login")
		v2.GET("/submit")
		v2.GET("/read")
	}

	router.Run(":3000")
}

func getting()  {

}

func posting()  {

}

func deleting()  {

}

func putting()  {

}

func patching()  {

}

func hesd()  {

}

func options()  {

}