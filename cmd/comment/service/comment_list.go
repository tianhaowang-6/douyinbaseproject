package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"
)

func CommentListService(ctx context.Context, req *comment.CommentListRequest) (cs []*comment.Comment,
	err error) {
	comments, err := db.CommentList(ctx, req.VideoId)
	cs = pack.Comments(comments)
	return cs, nil
}
