package db

import (
	"context"
	"douyin/cmd/comment/dal/model"
)

func CommentsCreate(ctx context.Context, comment *model.Comment) (err error) {
	return Q.WithContext(ctx).Comment.Create(comment)
}

func CommentList(ctx context.Context, videoId int64) (comments []*model.Comment, err error) {
	return Q.WithContext(ctx).Comment.Where(Q.Comment.VideoID.Eq(videoId)).Order(Q.Comment.CreatedAt.Desc()).Find()
}
