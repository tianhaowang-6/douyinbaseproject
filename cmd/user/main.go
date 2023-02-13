package main

import (
	"douyin-pro/cmd/user/dal/db"
	user "douyin-pro/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	db.Init()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr),
		server.WithServiceAddr(addr), // address
		//server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		//server.WithMuxTransport(),                                          // Multiplex
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
