package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/get/cookie", getCookie)
	router.GET("/set/cookie", setCookie)

	router.Run(":6240")

}

func getCookie(ctx *gin.Context) {

	cName := ctx.Query("cName")

	cookie, err := ctx.Cookie(cName)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": http.StatusOK})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "请求成功", "status": http.StatusOK, "data": cookie})
}

func setCookie(ctx *gin.Context) {

	cName := ctx.Query("cName")
	value := ctx.Query("value")
	ctx.SetCookie(cName, value, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"msg": "设置成功", "status": http.StatusOK})

}
