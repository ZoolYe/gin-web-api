package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {

		//当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in async ：" + cCp.Request.URL.Path)
		}()
		RequestOk(c)
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in sync ：" + c.Request.URL.Path)
		RequestOk(c)
	})

	r.Run(":6240")

}

func RequestOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK})
}
