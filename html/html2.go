package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	/*
		使用不同目录下名称相同的模板
	*/
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.Run(":6240")
}
