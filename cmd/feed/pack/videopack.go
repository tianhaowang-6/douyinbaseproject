package pack

import (
	"douyin/cmd/feed/dal/model"
	"douyin/kitex_gen/feed"
	"douyin/pkg/consts"
)

func Video(v *model.Video) (video feed.Video, err error) {
	if v == nil {
		return video, err
	}
	video = feed.Video{
		Id:            v.ID,
		Author:        nil,
		PlayUrl:       "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + v.PlayURL,
		CoverUrl:      "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    false,
		Title:         v.Title,
	}
	return video, nil
}

func Videos(vs []*model.Video) ([]*feed.Video, error) {
	videos := make([]*feed.Video, 0)
	for _, v := range vs {
		video := feed.Video{
			Id:            v.ID,
			Author:        nil,
			PlayUrl:       "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + v.PlayURL,
			CoverUrl:      "http://" + consts.WebServerPublicIP + ":" + consts.StaticPort + "/" + v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		}
		videos = append(videos, &video)
	}
	return videos, nil
}
