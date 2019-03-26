package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 获取路由参数
	router.GET("/user/:name",func(c *gin.Context){
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)
	})


	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + ":" + action
		c.String(http.StatusOK,message)
	})

	// 获取url查询参数
	router.GET("/welcome", func(c *gin.Context){
		firstname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK,"Hello %s %s",firstname,lastname)
	})

	// Multipart Urlencoded 表单
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick","anonymous")

		c.JSON(200,gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	// 获取post表单数据(url带查询参数)

	router.POST("/post_url", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s;name: %s;message: %s",id,page,name,message)
	})

	// 	映射参数 表单参数
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v;names: %v",ids,names)
	})

	// 文件上传
	router.POST("/upload_siglefile", func(c *gin.Context) {
		//single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)


		c.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!",file.Filename))
	})

	// 多文件上传
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _,file := range files{
			log.Println(file.Filename)
		}

		c.String(http.StatusOK,fmt.Sprintf("  %d files uploaded!",len(files)))
		})
	router.Run(":8080")
	}