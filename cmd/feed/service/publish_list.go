package service

import (
	"context"
	"douyin/cmd/feed/dal/db"
	"douyin/cmd/feed/dal/model"
	"douyin/kitex_gen/feed"
)

func PublishList(ctx context.Context, req *feed.PublishListRequest) (
	status_code int32, status_msg string, video_list []*model.Video, next_time int64, err error) {
	userId := req.GetUserId()
	videos, err := db.GetFeedListByUserId(ctx, userId)
	if err != nil {
		return 1, "查询失败", nil, 0, err
	}
	return 0, "", videos, 0, nil
}
