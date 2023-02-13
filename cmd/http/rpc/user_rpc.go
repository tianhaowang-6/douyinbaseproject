package rpc

import (
	"context"
	"douyin-pro/kitex_gen/user"
	"douyin-pro/kitex_gen/user/userservice"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

var userClient userservice.Client

func Init() {

	client, err := userservice.NewClient("userservice", client.WithHostPorts("0.0.0.0:8888"),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	userClient = client
}

func CreateUser(ctx context.Context, req user.DouyinUserRegisterRequest) error {
	println("rpc ... create")
	resp, err := userClient.CreateUser(ctx, &req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errors.New(*resp.StatusMsg)
	}
	println("创建成功")
	return nil
}

func UserLogin(ctx context.Context, req user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {

	resp, err = userClient.CheckUser(ctx, &req)
	fmt.Printf("resp=%#v\n", resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg)
	}
	println("登录成功")
	resp.StatusCode = 200
	return resp, nil
}

func GetUserInfo(ctx context.Context, req user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)
	respUserInfo, err := userClient.GetUserInfo(ctx, &req)
	if err != nil {
		return nil, err
	}
	if respUserInfo.StatusCode != 0 {
		return nil, errors.New(*resp.StatusMsg)
	}
	println("查询成功")
	resp.User = respUserInfo.User
	resp.StatusCode = 200
	return resp, nil
}
