package service

import (
	"context"
	"crypto/md5"
	"douyin-pro/cmd/user/dal/db"
	"douyin-pro/cmd/user/dal/model"
	"douyin-pro/kitex_gen/user"
	"douyin-pro/pkg/errno"
	"fmt"
	"io"
	"log"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.DouyinUserRegisterRequest) error {
	fmt.Printf("\nname ==%#v  -------\n", req.Username)
	users, err := db.QueryUser(s.ctx, req.Username)

	if err != nil {
		log.Fatal("error ", err)
		return err
	}

	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	//生成md5
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	println("db开始创建")
	return db.CreateUser(s.ctx, []*model.User{{Username: req.Username, Password: password}})
}
