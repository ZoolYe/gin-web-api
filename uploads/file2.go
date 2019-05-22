package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()

		files := form.File["file[]"]

		for _, file := range files {
			c.SaveUploadedFile(file, "C:/Users/ZOOL/Desktop/uploads/"+file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功", "status": http.StatusOK})
	})

	router.Run(":6240")
}
