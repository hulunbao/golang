package main

import (
	"github.com/gin-gonic/gin"
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

	router.Run(":8080")
}