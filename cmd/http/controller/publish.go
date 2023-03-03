package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/feed"
	"douyin/pkg/cover"
	"douyin/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 鉴权
	token := c.PostForm("token")
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't login",
		})
	}
	// 获取数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 保存数据
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%s", parseToken.Username, filename)
	// mp4文件名称
	saveFile := filepath.Join("./public/", finalName)
	// png文件名称
	saveCover := strings.ReplaceAll(saveFile, ".mp4", ".png")
	if err2 := c.SaveUploadedFile(data, saveFile); err2 != nil {
		log.Println(err2)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err2.Error(),
		})
		return
	}
	err = cover.VideoSaveToImg(saveFile, saveCover, 100)
	if err != nil {
		log.Println(err)
		return
	}
	req := feed.PublishActionRequest{
		Token:     token,
		FilePath:  saveFile,
		CoverPath: saveCover,
		Title:     c.Query("title"),
	}
	resp, err := rpc.PublishAction(context.Background(), &req)
	c.JSON(http.StatusOK, resp)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(userId, 10, 64)
	req := feed.PublishListRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.PublishList(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
