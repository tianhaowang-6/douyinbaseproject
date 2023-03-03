package service

import (
	"context"
	"douyin/cmd/feed/dal/db"
	"douyin/cmd/feed/dal/model"
	"douyin/kitex_gen/feed"
	"douyin/pkg/jwt"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func PublishAction(ctx context.Context, req *feed.PublishActionRequest) (
	status_code int32, status_msg string, err error) {
	token := req.Token
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		return 1001, "鉴权失败", err
	}
	video := model.Video{
		ID:            0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
		AuthorID:      parseToken.UserID,
		PlayURL:       req.FilePath,
		CoverURL:      req.CoverPath,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	videos := make([]*model.Video, 0)
	videos = append(videos, &video)
	fmt.Printf("%#v\n", video)
	db.PublishAction(ctx, videos)
	return 0, "发布成功", nil
}
