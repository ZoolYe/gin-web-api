package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	r1 := router.Group("/v1")
	{
		r1.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "我是v1的login", "status": http.StatusOK})
		})
	}

	r2 := router.Group("/v2")
	{
		r2.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "我是v2的login", "status": http.StatusOK})
		})
	}

	router.Run(":6240")

}
