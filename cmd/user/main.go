package main

import (
	"douyin/cmd/user/dal/db"
	user "douyin/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	//r, err := etcd.NewEtcdRegistry([]string{"192.168.200.131:2379"})
	db.Init()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))
	//server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userservice"}),
	////server.WithRegistry(r))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
