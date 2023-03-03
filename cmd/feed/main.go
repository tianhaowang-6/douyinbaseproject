package main

import (
	"douyin/cmd/feed/dal/db"
	feed1 "douyin/kitex_gen/feed/feedservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	db.Init()
	addr, _ := net.ResolveTCPAddr("tcp", consts.FeedServiceIPPORT)
	svr := feed1.NewServer(new(FeedServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
