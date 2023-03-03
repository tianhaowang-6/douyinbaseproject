package db

import (
	"context"
	"douyin/cmd/feed/dal/model"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_PublicAction(t *testing.T) {
	Init()
	videos := make([]*model.Video, 0)
	videos = append(videos, &model.Video{
		ID:            0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
		AuthorID:      1,
		PlayURL:       "aa",
		CoverURL:      "bb",
		FavoriteCount: 0,
		CommentCount:  0,
	})
	err := PublishAction(context.Background(), videos)
	if err != nil {
		return
	}
	return
}
