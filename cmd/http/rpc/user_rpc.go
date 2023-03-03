package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

var userClient userservice.Client

func InitUser() {
	//r, err := etcd.NewEtcdResolver([]string{"192.168.200.131:2379"})
	client, err := userservice.NewClient("userservice",
		client.WithHostPorts("127.0.0.1:8888"), //服务地址
		//client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	userClient = client
}

func CreateUser(ctx context.Context, req user.UserRegisterRequest) error {
	// rpc 创建用户
	r, err := userClient.UserRegister(ctx, &req)
	fmt.Printf("?%#v,err=%#v", r, err)
	if err != nil {
		return err
	}
	return nil
}

func UserLogin(ctx context.Context, req user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp, err = userClient.UserLogin(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg)
	}
	println("登录成功")
	resp.StatusCode = 0
	return resp, nil
}

func GetUserInfo(ctx context.Context, req user.UserRequest) (resp *user.UserResponse, err error) {
	respUserInfo, err := userClient.UserInfo(ctx, &req)
	fmt.Printf("%#v", respUserInfo.BaseReponse)
	if err != nil {
		return nil, err
	}
	if respUserInfo.BaseReponse.StatusCode != 0 {
		return nil, errors.New(respUserInfo.String())
	}
	return respUserInfo, nil
}
