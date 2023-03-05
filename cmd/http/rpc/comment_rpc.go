package rpc

import (
	"context"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/comment/commentservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

var commentClient commentservice.Client

func InitComment() {
	//r, err := etcd.NewEtcdResolver([]string{"192.168.200.131:2379"})
	client, err := commentservice.NewClient(consts.CommentServiceName,
		client.WithHostPorts(consts.CommentServiceIPPORT), //服务地址
		//client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	commentClient = client
}

// CommentAction implements the CommentServiceImpl interface.
func CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	resp, err = commentClient.CommentAction(ctx, req)
	return resp, err
}

// CommentList implements the CommentServiceImpl interface.
func CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
