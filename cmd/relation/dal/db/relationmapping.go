package db

import (
	"context"
	"douyin/cmd/relation/dal/model"
)

func CreateFollow(ctx context.Context, f *model.Follow) error {
	return Q.WithContext(ctx).Follow.Create(f)
}

func FollowFirstOrCreate(ctx context.Context, f *model.Follow) (rf *model.Follow, err error) {
	return Q.WithContext(ctx).Follow.Where(Q.Follow.UserID.Eq(f.UserID)).Where(Q.Follow.ToUserID.Eq(f.ToUserID)).FirstOrCreate()
}
