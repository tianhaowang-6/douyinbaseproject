package pack

import (
	"douyin/cmd/user/dal/model"
	"douyin/kitex_gen/user"
)

func UserInfo(u *model.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
