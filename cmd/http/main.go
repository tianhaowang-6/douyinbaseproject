package main

import (
	"douyin/cmd/http/rpc"
	"douyin/cmd/http/service"
	"github.com/gin-gonic/gin"
)

func main() {
	rpc.Init()
	go service.RunMessageServer()
	r := gin.Default()
	initRouter(r)
	r.Run()
}
