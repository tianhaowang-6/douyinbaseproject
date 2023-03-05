package pack

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/model"
	"douyin/kitex_gen/comment"
)

func Comment(commen *model.Comment) (c *comment.Comment) {
	userID := commen.UserID
	users, _ := db.QueryUser(context.Background(), string(userID))
	c2 := comment.Comment{
		Id:         commen.ID,
		User:       User(users[0]),
		Content:    commen.Content,
		CreateDate: "",
		VideoId:    commen.VideoID,
	}
	return &c2
}

func Comments(commens []*model.Comment) (cs []*comment.Comment) {
	comments := make([]*comment.Comment, 0)
	for _, commen := range commens {
		c := Comment(commen)
		comments = append(comments, c)
	}
	return comments
}
func User(user *model.User) (u *comment.User) {
	return &comment.User{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   &user.FollowCount,
		FollowerCount: &user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
