package db

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/dal/model"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	Init()
	users := make([]*model.User, 0)

	h := md5.New()
	password := "qwer1234"
	io.WriteString(h, password)
	password = fmt.Sprintf("%x", h.Sum(nil))
	users = append(users, &model.User{
		ID:            0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		Username:      "小明",
		Password:      password,
	})
	CreateUser(context.Background(), users)
}

func TestQueryUser(t *testing.T) {
	Init()
	user, err := QueryUser(context.Background(), "小红")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user[0])
}

func TestMGetUsers(t *testing.T) {
	Init()
	users, err := MGetUsers(context.Background(), []int64{1, 2, 3, 4, 5, 6, 7, 9, 10})
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		log.Println(user)
	}
}
