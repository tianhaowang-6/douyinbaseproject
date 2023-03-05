package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/relation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	request := relation.RelationActionRequest{
		Token:      token,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	}
	resp, _ := rpc.RelationAction(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	request := relation.RelationFollowListRequest{
		UserId: userId,
		Token:  token,
	}
	resp, _ := rpc.RelationFollowList(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {

	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	request := relation.RelationFollowerListRequest{
		UserId: userId,
		Token:  token,
	}
	resp, _ := rpc.RelationFollowerList(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	request := relation.RelationFriendListRequest{
		UserId: userId,
		Token:  token,
	}
	resp, _ := rpc.RelationFriendList(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}
