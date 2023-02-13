package pack

import (
	"douyin-pro/cmd/user/dal/model"
	"douyin-pro/kitex_gen/user"
)

func UserInfo(u *model.UserInfo) *user.UserInfo {
	if u == nil {
		return nil
	}
	return &user.UserInfo{
		Id:            u.ID,
		Name:          u.Name,
		FollowCount:   &u.FollowCount,
		FollowerCount: &u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
