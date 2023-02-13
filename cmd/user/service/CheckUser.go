package service

import (
	"context"
	"crypto/md5"
	"douyin-pro/cmd/user/dal/db"
	"douyin-pro/kitex_gen/user"
	"douyin-pro/pkg/errno"
	"fmt"
	"io"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) (c *CheckUserService) {
	return &CheckUserService{ctx}
}

func (s *CheckUserService) CheckUser(req *user.DouyinUserLoginRequest) (id int64, err error) {
	h := md5.New()

	_, err = io.WriteString(h, req.Password)
	if err != nil {
		return 0, nil
	}
	username := req.Username
	users, err := db.QueryUser(s.ctx, username)
	if err != nil {
		return 0, nil
	}
	user := users[0]
	password := fmt.Sprintf("%x", h.Sum(nil))
	println("userpasswor=", password, "dbpassword=", user.Password, "len=", len(users), "\n")
	if len(users) != 1 || user.Password != password {
		return 0, errno.AuthorizationFailedErr
	}
	return user.ID, nil
}
