package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		log.Printf("ids: %v ; names: %v", ids, names)
	})

	router.Run(":6240")
}
