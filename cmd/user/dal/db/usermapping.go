package db

import (
	"context"
	"douyin/cmd/user/dal/model"
)

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	return Q.WithContext(ctx).User.Where(Q.User.ID.In(userIDs...)).Find()
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*model.User) error {
	return Q.WithContext(ctx).User.Create(users...)
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*model.User, error) {
	return Q.WithContext(ctx).User.Where(Q.User.Username.Eq(userName)).Find()
}
func QueryUserById(ctx context.Context, id int64) ([]*model.User, error) {
	return Q.WithContext(ctx).User.Where(Q.User.ID.Eq(id)).Find()
}
