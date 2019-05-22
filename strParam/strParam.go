package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "nil")
		age := c.Query("age") //c.Request.URL.Query().Get("lastname") 的一种快捷方式
		c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": gin.H{"name": name, "age": age}})
	})
	router.Run(":6240")
}
