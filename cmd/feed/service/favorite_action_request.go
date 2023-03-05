package service

import (
	"context"
	"douyin/cmd/feed/dal/db"
	"douyin/cmd/feed/dal/model"
	"douyin/kitex_gen/feed"
	"douyin/pkg/jwt"
	"gorm.io/gorm"
	"time"
)

func FavoriteActionRequestService(ctx context.Context, req *feed.FavoriteActionRequest) (
	status_code int32, status_msg *string, err error) {
	token := req.GetToken()
	_, err = jwt.ParseToken(token)
	if err != nil {
		status_code = 1008
		return status_code, nil, err
	}
	claims, err := jwt.ParseToken(token)
	userID := claims.UserID
	videoId := req.VideoId
	favorite := model.Favorite{
		ID:        0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
		UserID:    userID,
		VideoID:   videoId,
	}
	fs := make([]*model.Favorite, 0)
	fs = append(fs, &favorite)
	dfs, err := db.FavoriteActionSelectByUserIdAndVideoId(ctx, userID, videoId)
	if len(dfs) != 0 {
		println("1111111111111111111")
		_, _ = db.FavoriteActionUpdate(ctx, dfs[0].ID, req.ActionType)
	} else {
		println("000000000000000000")
		err = db.FavoriteActionCreate(ctx, fs)
	}

	if err != nil {
		return 1009, nil, err
	}
	return 0, nil, nil
}
