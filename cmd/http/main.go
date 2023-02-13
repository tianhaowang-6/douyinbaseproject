package main

import (
	"douyin-pro/cmd/http/rpc"
	"douyin-pro/cmd/http/service"
	"github.com/gin-gonic/gin"
)

//func main() {
//	r := gin.Default()
//	r.GET("/", func(context *gin.Context) {
//		context.JSON(200, "hello")
//	})
//	r.Run()
//}

func main() {
	rpc.Init()
	go service.RunMessageServer()
	r := gin.Default()
	initRouter(r)
	r.GET("abcd", func(context *gin.Context) {
		context.JSON(200, "hello")
	})
	r.Run()
}
