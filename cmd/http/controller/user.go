package controller

import (
	"context"
	"douyin/cmd/http/rpc"
	"douyin/kitex_gen/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=
var usersLoginInfo = map[string]User{
	"zhanglei ": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	request := user.UserRegisterRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	//token := username + password
	println("rpc")
	err := rpc.CreateUser(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}

func Login(c *gin.Context) {
	println("login")
	request := user.UserLoginRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	resp, err := rpc.UserLogin(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "登录失败"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   resp.UserId,
			Token:    resp.Token,
		})
	}
}

func UserInfo(c *gin.Context) {
	//token := c.Query("token")
	userId := c.Query("user_id")
	token := c.Query("token")
	id, err := strconv.ParseInt(userId, 10, 64)
	req := user.UserRequest{
		UserId: id,
		Token:  token,
	}
	resp, err := rpc.GetUserInfo(context.Background(), req)
	userInfo := resp.User
	if err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User: User{
				Id:            userInfo.Id,
				Name:          userInfo.Name,
				FollowCount:   *userInfo.FollowCount,
				FollowerCount: *userInfo.FollowerCount,
				IsFollow:      userInfo.IsFollow,
			},
		})
	}
}
