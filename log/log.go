package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {

	file, _ := os.Create("C:/Users/ZOOL/Desktop/uploads/gin.log")
	//将日志写入到文件，不打印在控制台
	//gin.DefaultWriter = io.MultiWriter(file)
	//将日志写入到文件，并且打印在控制台
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK})
	})

	router.Run(":6240")

}
