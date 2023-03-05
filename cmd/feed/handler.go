package main

import (
	"context"
	"douyin/cmd/feed/pack"
	"douyin/cmd/feed/service"
	feed "douyin/kitex_gen/feed"
	"douyin/kitex_gen/user"
	"fmt"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	println("getFeed...")
	resp = new(feed.FeedResponse)
	code, msg, listV, nextTime, err := service.GetFeedService(ctx, req)
	fmt.Printf("-----%#v\n", code)
	if err != nil {
		resp.StatusCode = 1
		return resp, err
	}
	videos, err := pack.Videos(listV)
	// 封装user中的Author
	for i, video := range listV {
		_, _, u, _ := service.UserInfoService(ctx, &user.UserRequest{
			UserId: video.AuthorID,
			Token:  "",
		})
		if u == nil {
			continue
		}
		u2 := feed.User{
			Id:            video.ID,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      false,
		}
		videos[i].Author = &u2
	}
	if err != nil {
		return nil, err
	}
	resp.VideoList = videos
	resp.StatusMsg = &msg
	resp.StatusCode = code
	resp.NextTime = &nextTime
	fmt.Printf("%#v--\n", resp)
	return resp, nil
}

// PublishAction implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) PublishAction(ctx context.Context,
	req *feed.PublishActionRequest) (resp *feed.PublishActionResponse,
	err error) {
	// TODO: Your code here...
	println("PublishAction")
	code, msg, err := service.PublishAction(ctx, req)
	resp = new(feed.PublishActionResponse)
	resp.StatusCode = code
	resp.StatusMsg = &msg
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PublishList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) PublishList(ctx context.Context, req *feed.PublishListRequest) (resp *feed.PublishListResponse, err error) {
	// TODO: Your code here...
	println("PublishList")
	_, msg, listV, _, err := service.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	videos, err := pack.Videos(listV)
	for i, video := range listV {
		_, _, u, _ := service.UserInfoService(ctx, &user.UserRequest{
			UserId: video.AuthorID,
			Token:  "",
		})
		if u == nil {
			continue
		}
		u2 := feed.User{
			Id:            video.ID,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      false,
		}
		videos[i].Author = &u2
	}
	return &feed.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  videos,
	}, nil
}

// FavoriteAction implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) FavoriteAction(ctx context.Context, req *feed.FavoriteActionRequest) (resp *feed.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	println("FavoriteAction")
	code, msg, err := service.FavoriteActionRequestService(ctx, req)
	response := feed.FavoriteActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
	return &response, nil
}

// FavoriteList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) FavoriteList(ctx context.Context, req *feed.FavoriteListRequest) (resp *feed.FavoriteListResponse, err error) {
	// TODO: Your code here...
	println("FavoriteList")
	favoriteList, err := service.FavoriteListService(ctx, req)
	if err != nil {
		return nil, err
	}
	response := feed.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  favoriteList,
	}
	return &response, nil
}
