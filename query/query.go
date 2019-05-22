package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		//参数在url上
		id := c.Query("id")
		//参数在url上，如果获取不到就给默认值
		page := c.DefaultQuery("page", "0")
		//form-data传参
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.Run(":6240")
}
