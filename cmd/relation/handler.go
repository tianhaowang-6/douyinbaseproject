package main

import (
	"context"
	"douyin/cmd/relation/service"
	relation "douyin/kitex_gen/relation"
	"douyin/pkg/jwt"
	"errors"
)

// RelationShipServiceImpl implements the last service interface defined in the IDL.
type RelationShipServiceImpl struct{}

// RelationAction implements the RelationShipServiceImpl interface.
func (s *RelationShipServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	token := req.Token
	_, err = jwt.ParseToken(token)
	if err != nil {
		return nil, errors.New("鉴权失败")
	}
	service.RelationAction(ctx, req)
	return
}

// RelationFollowList implements the RelationShipServiceImpl interface.
func (s *RelationShipServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationShipServiceImpl interface.
func (s *RelationShipServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFriendList implements the RelationShipServiceImpl interface.
func (s *RelationShipServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
