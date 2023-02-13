package db

import (
	"context"
	"douyin-pro/cmd/user/dal/model"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Init()
	users := make([]*model.User, 0)
	users = append(users, &model.User{Username: "wpctest", Password: "123187"})
	CreateUser(context.Background(), users)
}

func TestQueryUser(t *testing.T) {
	Init()
	user, err := QueryUser(context.Background(), "wpctest")
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

func TestQueryUserInfo(t *testing.T) {
	Init()
	userInfo, err := QueryUserInfo(context.Background(), 5)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userInfo)
}
