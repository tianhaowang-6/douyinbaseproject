package main

import (
	"douyin/cmd/feed/dal/db"
	message "douyin/kitex_gen/message/messageservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	db.Init()
	addr, _ := net.ResolveTCPAddr("tcp", consts.MessageServiceIPPORT)
	svr := message.NewServer(new(MessageServiceImpl), server.WithServiceAddr(addr))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
