package service

import (
	"context"
	"douyin/cmd/feed/dal/db"
	"douyin/cmd/feed/dal/model"
	"douyin/kitex_gen/feed"
	"time"
)

func GetFeedService(ctx context.Context, req *feed.FeedRequest) (
	status_code int32, status_msg string, video_list []*model.Video, next_time int64, err error) {

	videos, err := db.GetFeedListByLeastTime(ctx, req.LatestTime)
	if err != nil {
		return 1001, "查询视频失败", nil, 0, err
	}
	var nextTime int64
	if len(videos) == 0 {
		nextTime = time.Now().Unix() - 1
	} else {
		nextTime = videos[len(videos)-1].CreatedAt.Unix() - 1
	}

	return 0, "", videos, nextTime, nil
}
