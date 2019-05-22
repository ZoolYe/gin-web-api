package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {

		objA := formA{}
		objB := formB{}

		// c.ShouldBind 使用了 c.Request.Body，不可重用。
		if errA := c.ShouldBind(&objA); errA == nil {
			log.Println(objA)
			c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": &objA})
			// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		} else if errB := c.ShouldBind(&objB); errB == nil {
			log.Println(objB)
			c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": &objB})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "失败", "status": http.StatusInternalServerError})
		}
	})

	r.POST("/v2", func(c *gin.Context) {

		objA := formA{}
		objB := formB{}

		// 读取 c.Request.Body 并将结果存入上下文。
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			log.Println(objA)
			c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": &objA})
			// 这时, 复用存储在上下文中的 body。
		} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": &objB})
			// 可以接受其他格式
		} else if errXml := c.ShouldBindBodyWith(&objB, binding.XML); errXml == nil {
			c.JSON(http.StatusOK, gin.H{"msg": "成功", "status": http.StatusOK, "data": &objB})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "失败", "status": http.StatusInternalServerError})
		}
	})

	r.Run(":6240")
}
