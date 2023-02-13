package service

import (
	"context"
	"douyin-pro/cmd/user/dal/db"
	"douyin-pro/cmd/user/pack"
	"douyin-pro/kitex_gen/user"
	"fmt"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) (c *QueryUserService) {
	return &QueryUserService{ctx}
}

func (s *QueryUserService) QueryUserInfo(req *user.DouyinUserRequest) (userInfo *user.UserInfo, err error) {
	println("userId====", req.UserId)
	userInfos, err := db.QueryUserInfo(s.ctx, req.UserId)
	fmt.Printf("userinfos===%#v\n", userInfos)
	if err != nil {
		return nil, err
	}
	info := userInfos[0]
	fmt.Printf("userinfo===%#v\n", info)
	return pack.UserInfo(info), nil

}
