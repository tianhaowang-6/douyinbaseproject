package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationshipservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

var relationClient relationshipservice.Client

func InitRelation() {
	//r, err := etcd.NewEtcdResolver([]string{"192.168.200.131:2379"})
	client, err := relationshipservice.NewClient(consts.RelationServiceName,
		client.WithHostPorts(consts.RelationServiceIPPORT), //服务地址
		//client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	relationClient = client
}

// RelationAction implements the RelationShipServiceImpl interface.
func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the RelationShipServiceImpl interface.
func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationShipServiceImpl interface.
func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFriendList implements the RelationShipServiceImpl interface.
func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
