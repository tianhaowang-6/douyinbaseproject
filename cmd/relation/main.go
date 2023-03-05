package main

import (
	"douyin/cmd/feed/dal/db"
	relation "douyin/kitex_gen/relation/relationshipservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	db.Init()
	addr, _ := net.ResolveTCPAddr("tcp", consts.RelationServiceIPPORT)
	svr := relation.NewServer(new(RelationShipServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
