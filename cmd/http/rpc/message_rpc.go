package rpc

import (
	"context"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messageservice"
	"douyin/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

var messageClient messageservice.Client

func InitMessage() {
	//r, err := etcd.NewEtcdResolver([]string{"192.168.200.131:2379"})
	client, err := messageservice.NewClient(consts.MessageServiceName,
		client.WithHostPorts(consts.MessageServiceIPPORT), //服务地址
		//client.WithResolver(r),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}
	messageClient = client
}

// MessageChat implements the MessageServiceImpl interface.
func MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationActionMessage implements the MessageServiceImpl interface.
func RelationActionMessage(ctx context.Context, req *message.RelationActionMessageRequest) (resp *message.RelationActionMessageResponse, err error) {
	// TODO: Your code here...
	return
}
