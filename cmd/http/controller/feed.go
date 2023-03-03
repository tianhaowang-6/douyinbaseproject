package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/feed"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	token := c.Query("token")
	feedRequest := feed.FeedRequest{
		LatestTime: &latestTime,
		Token:      &token,
	}
	resp, err := rpc.GetFeed(context.Background(), &feedRequest)
	fmt.Printf("\nresp.list=%#v\n", resp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "",
			},
			VideoList: nil,
			NextTime:  0,
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
