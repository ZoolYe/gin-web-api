package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//HTML 渲染
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "我是tmpl模板",
		})
	})

	router.Run(":6240")
}
