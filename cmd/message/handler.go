package main

import (
	"context"
	message "douyin/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationActionMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) RelationActionMessage(ctx context.Context, req *message.RelationActionMessageRequest) (resp *message.RelationActionMessageResponse, err error) {
	// TODO: Your code here...
	return
}
