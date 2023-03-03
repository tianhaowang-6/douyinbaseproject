package db

import (
	"context"
	"douyin/cmd/feed/dal/model"
	"time"
)

func PublishAction(ctx context.Context, videos []*model.Video) error {
	return Q.WithContext(ctx).Video.Create(videos...)
}

func GetFeedListByLeastTime(ctx context.Context, latestTime *int64) (videos []*model.Video, err error) {
	var curTime time.Time
	if latestTime == nil || *latestTime == 0 || *latestTime > 1076997151921 {
		curTime = time.Now()
	} else {
		curTime = time.Unix(*latestTime, 0)
	}
	curTime = curTime.Add(8 * time.Hour)
	const limit = 30
	return Q.WithContext(ctx).Video.Where(Q.Video.CreatedAt.Lte(curTime)).Limit(limit).
		Order(Q.Video.CreatedAt.Desc()).Find()
}
func GetFeedListByUserId(ctx context.Context, userId int64) (videos []*model.Video, err error) {
	return Q.WithContext(ctx).Video.Where(Q.Video.AuthorID.Eq(userId)).Limit(30).Order(Q.Video.CreatedAt.Desc()).Find()
}
