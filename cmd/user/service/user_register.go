package service

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/dal/model"
	"douyin/kitex_gen/user"
	"fmt"
	"gorm.io/gorm"
	"io"
	"time"
)

func UserRegisterService(
	ctx context.Context, req *user.UserRegisterRequest) (
	status_code int32, status_msg string, user_id int64, token string, err error) {

	// 判断昵称是否合格
	qUsers, err := db.QueryUser(ctx, req.Username)
	if err != nil || len(qUsers) != 0 {
		return 1001, "已经有用户创建失败", 0, "", err
	}

	// 创建
	h := md5.New()
	_, err = io.WriteString(h, req.Password)
	password := fmt.Sprintf("%x", h.Sum(nil))
	if err != nil {

	}
	users := make([]*model.User, 0)
	users = append(users, &model.User{
		ID:            0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		Username:      req.Username,
		Password:      password,
	})
	err = db.CreateUser(ctx, users)
	if err != nil {
		return 1003, "", 0, "", err
	}
	return 0, "成功", 0, "", nil
}
