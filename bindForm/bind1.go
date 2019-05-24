package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Login struct {
	//一个字段的 tag 加上了 binding:"required"，但绑定时是空值, Gin 会报错
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {

	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var login Login
		/*
			如果发生绑定错误，则请求终止，并触发 c.AbortWithError(400, err).SetType(ErrorTypeBind)
			响应状态码被设置为 400
		*/
		c.MustBindWith(&login, binding.JSON)

		c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "status": http.StatusOK, "data": login})
	})

	router.POST("/v2/login", func(c *gin.Context) {
		var login Login

		//如果发生绑定错误，Gin 会返回错误并由开发者处理错误和请求
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "请求失败",
				"status": http.StatusInternalServerError, "data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "status": http.StatusOK, "data": login})
	})

	router.Run(":6240")
}
