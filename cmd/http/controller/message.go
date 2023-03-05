package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/message"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	content := c.Query("content")
	request := message.RelationActionMessageRequest{
		Token:      token,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
		Content:    content,
	}
	resp, _ := rpc.RelationActionMessage(context.Background(), &request)
	c.JSON(http.StatusOK, resp)
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {

	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	preMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	request := message.MessageChatRequest{
		Token:      token,
		ToUserId:   toUserId,
		PreMsgTime: preMsgTime,
	}
	resp, _ := rpc.MessageChat(context.Background(), &request)
	c.JSON(http.StatusOK, resp)

}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
