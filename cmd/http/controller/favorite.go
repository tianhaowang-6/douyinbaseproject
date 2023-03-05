package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/feed"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	request := feed.FavoriteActionRequest{
		Token:      token,
		VideoId:    videoId,
		ActionType: int32(actionType),
	}
	resp, err := rpc.FavoriteAction(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "gg"})
	}
	c.JSON(http.StatusOK, resp)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	request := feed.FavoriteListRequest{
		UserId: userId,
		Token:  token,
	}
	resp, _ := rpc.FavoriteList(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}
