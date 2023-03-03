package rpc

import (
	"context"
	"douyin/kitex_gen/feed"
	"douyin/kitex_gen/feed/feedservice"
	"douyin/pkg/consts"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

var feedClient feedservice.Client

func InitFeed() {
	//r, err := etcd.NewEtcdResolver([]string{"192.168.200.131:2379"})
	client, err := feedservice.NewClient(consts.FeedServiceName,
		client.WithHostPorts(consts.FeedServiceIPPORT), //服务地址
		//client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	feedClient = client
}

// GetFeed implements the FeedServiceImpl interface.
func GetFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	response, err := feedClient.GetFeed(ctx, req)
	fmt.Printf("%#v", resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// PublishAction implements the FeedServiceImpl interface.
func PublishAction(ctx context.Context, req *feed.PublishActionRequest) (resp *feed.PublishActionResponse, err error) {
	// TODO: Your code here...
	response, err := feedClient.PublishAction(ctx, req)
	return response, err
}

// PublishList implements the FeedServiceImpl interface.
func PublishList(ctx context.Context, req *feed.PublishListRequest) (resp *feed.PublishListResponse, err error) {
	// TODO: Your code here...
	r, err := feedClient.PublishList(ctx, req)
	return r, err
}

// FavoriteAction implements the FeedServiceImpl interface.
func FavoriteAction(ctx context.Context, req *feed.FavoriteActionRequest) (resp *feed.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	r, err := feedClient.FavoriteAction(ctx, req)
	return r, err
}

// FavoriteList implements the FeedServiceImpl interface.
func FavoriteList(ctx context.Context, req *feed.FavoriteListRequest) (resp *feed.FavoriteListResponse, err error) {
	// TODO: Your code here...
	r, err := feedClient.FavoriteList(ctx, req)
	return r, err
}
