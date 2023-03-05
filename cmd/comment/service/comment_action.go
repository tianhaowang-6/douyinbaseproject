package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/dal/model"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/comment"
	"douyin/pkg/jwt"
	"gorm.io/gorm"
	"time"
)

func CommentActionService(ctx context.Context, req *comment.CommentActionRequest) (statusCode int64,
	statusMsg *string, comment *comment.Comment, err error) {
	token := req.Token
	parseToken, err := jwt.ParseToken(token)
	id := parseToken.UserID
	comme := model.Comment{
		ID:        0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
		UserID:    id,
		Content:   *req.CommentText,
	}
	err = db.CommentsCreate(ctx, &comme)
	println("commentId", comme.ID)
	if err != nil {
		return 1, nil, nil, nil
	}
	return 0, nil, pack.Comment(&comme), nil
}
