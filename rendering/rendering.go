package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "成功", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {

		var msg struct {
			// 注意 msg.Name 在 JSON 中变成了 "user"
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "zoolye"
		msg.Message = "成功"
		msg.Number = 134
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "成功", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "成功", "status": http.StatusOK})
	})

	r.Run(":6240")
}
