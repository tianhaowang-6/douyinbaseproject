package main

import (
	"context"
	"douyin/cmd/comment/service"
	comment "douyin/kitex_gen/comment"
	"douyin/pkg/jwt"
	"errors"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	token := req.Token
	_, err = jwt.ParseToken(token)
	if err != nil {
		return nil, errors.New("鉴权失败")
	}
	response := new(comment.CommentActionResponse)
	code, msg, c, err := service.CommentActionService(ctx, req)
	response.Comment = c
	response.StatusCode = int32(code)
	response.StatusMsg = msg
	return response, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	token := req.GetToken()
	_, err = jwt.ParseToken(token)
	if err != nil {
		return nil, errors.New("鉴权失败")
	}
	cs, err := service.CommentListService(ctx, req)
	return &comment.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   nil,
		CommentList: cs,
	}, nil
}
