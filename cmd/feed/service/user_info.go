package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/jwt"
)

func UserInfoService(
	ctx context.Context, req *user.UserRequest) (
	status_code int32, status_msg string, user *user.User, err error) {
	//鉴权
	token := req.Token
	// 解析发起者用户权限
	_, err = jwt.ParseToken(token)
	if err != nil {
		status_code = 1008
		return status_code, "没有通过鉴权", nil, err
	}
	// 查询
	users, err := db.QueryUserById(ctx, req.UserId)
	if err != nil || len(users) == 0 {
		return 1001, "gg", user, err
	}

	return 0, "success", pack.UserInfo(users[0]), nil
}
