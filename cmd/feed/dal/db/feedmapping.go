package db

import (
	"context"
	"douyin/cmd/feed/dal/model"
	"gorm.io/gen"
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
func GetFeedByVideoId(ctx context.Context, videoId int64) (videos *model.Video, err error) {
	return Q.WithContext(ctx).Video.Where(Q.Video.ID.Eq(videoId)).First()
}
func FavoriteActionCreate(ctx context.Context, favorites []*model.Favorite) (err error) {
	return Q.WithContext(ctx).Favorite.Create(favorites...)
}

func FavoriteActionUpdate(ctx context.Context, id int64, typ int32) (info gen.ResultInfo, err error) {
	return Q.WithContext(ctx).Favorite.Where(Q.Favorite.ID.Eq(id)).Update(Q.Favorite.ActionType, typ)
}
func FavoriteActionSelectByUserIdAndVideoId(ctx context.Context, userId int64, videoId int64) (fs []*model.Favorite, err error) {
	return Q.WithContext(ctx).Favorite.Where(Q.Favorite.UserID.Eq(userId)).Where(Q.Favorite.VideoID.Eq(videoId)).Find()
}
func FavoriteListByUserId(ctx context.Context, userId int64) ([]*model.Favorite, error) {
	return Q.WithContext(ctx).Favorite.Where(Q.Favorite.UserID.Eq(userId)).Order(Q.Favorite.CreatedAt).Find()
}
