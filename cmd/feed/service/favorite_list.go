package service

import (
	"context"
	"douyin/cmd/feed/dal/db"
	"douyin/cmd/feed/pack"
	"douyin/kitex_gen/feed"
	"fmt"
)

func FavoriteListService(ctx context.Context, req *feed.FavoriteListRequest) (
	favoriteList []*feed.Video, err error) {
	userId := req.GetUserId()
	println(userId, "userId")
	favorites, err := db.FavoriteListByUserId(ctx, userId)
	println("favoritesLen=", len(favorites))
	videos := make([]*feed.Video, 0)
	for _, favorite := range favorites {
		video, _ := db.GetFeedByVideoId(ctx, favorite.VideoID)
		packV, _ := pack.Video(video)
		fmt.Printf("packV:%#v\n", packV)
		videos = append(videos, &packV)
	}
	return videos, nil
}
