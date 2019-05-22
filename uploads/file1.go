package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 //8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		errs := c.SaveUploadedFile(file, "C:/Users/ZOOL/Desktop/uploads/"+file.Filename)
		if errs != nil {
			log.Println(errs.Error())
			c.JSON(http.StatusInternalServerError, errs.Error())
			return
		}
		c.JSON(http.StatusOK, "上传成功")
	})

	router.Run(":6240")
}
