package service

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/jwt"
	"fmt"
	"io"
)

func UserLoginService(
	ctx context.Context, req *user.UserLoginRequest) (
	status_code int32, status_msg string, user_id int64, token string, err error) {
	username := req.Username
	password := req.Password
	if username == "" || password == "" {
		return 1001, "登录失败", 0, "", err
	}
	users, err := db.QueryUser(ctx, username)
	if len(users) != 1 {
		return 1002, "没有该用户", 0, "", err
	}
	h := md5.New()
	_, err = io.WriteString(h, req.Password)
	password = fmt.Sprintf("%x", h.Sum(nil))
	if password != users[0].Password {
		return 1003, "登录失败", 0, "", err
	}
	genToken, err := jwt.GenToken(users[0].ID, username)
	return 0, "登录成功", users[0].ID, genToken, err
}
