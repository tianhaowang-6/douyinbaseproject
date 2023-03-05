package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/model"
	"douyin/kitex_gen/relation"
	"douyin/pkg/jwt"
	"gorm.io/gorm"
	"time"
)

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) error {
	token := req.GetToken()
	parseToken, _ := jwt.ParseToken(token)
	userID := parseToken.UserID
	follow := model.Follow{
		ID:         0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  gorm.DeletedAt{},
		UserID:     userID,
		ToUserID:   req.ToUserId,
		ActionType: req.ActionType,
	}
	_, err := db.FollowFirstOrCreate(ctx, &follow)
	return err
}
