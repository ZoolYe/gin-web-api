package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {

	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":6240")
}

func startPage(c *gin.Context) {
	var person Person
	//ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "Success")
}
