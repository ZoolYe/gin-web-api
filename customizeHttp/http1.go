package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	//自定义http配置
	http1()
	http2()
}

func http1() {

	router := gin.Default()
	http.ListenAndServe(":6240", router)

}

func http2() {

	router := gin.Default()

	s := &http.Server{
		Addr:           ":6240",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
