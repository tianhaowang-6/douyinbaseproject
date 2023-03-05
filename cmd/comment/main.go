package main

import (
	"douyin/cmd/feed/dal/db"
	comment "douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	db.Init()
	addr, _ := net.ResolveTCPAddr("tcp", consts.CommentServiceIPPORT)
	svr := comment.NewServer(new(CommentServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
