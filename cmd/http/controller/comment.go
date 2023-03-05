package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	commentType := c.Query("comment_type")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	request := comment.CommentActionRequest{
		Token:       token,
		VideoId:     videoId,
		ActionType:  int32(actionType),
		CommentText: &commentType,
		CommentId:   &commentId,
	}
	resp, err := rpc.CommentAction(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	request := comment.CommentListRequest{
		Token:   token,
		VideoId: videoId,
	}
	resp, err := rpc.CommentList(context.Background(), &request)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
