package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}
		//使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/zoolye", func(c *gin.Context) {
		data := map[string]interface{}{
			"web:":   "www.zoolye.com",
			"email:": "www.iphone@foxmail.com",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":6240")
}
