package main

import (
	"context"
	"douyin/cmd/user/service"
	user "douyin/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(
	ctx context.Context, req *user.UserRegisterRequest) (
	resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserRegisterResponse)
	code, msg, id, token, err := service.UserRegisterService(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.UserId = id
	resp.StatusCode = code
	resp.StatusMsg = &msg
	resp.Token = token
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.UserLoginResponse)
	code, msg, id, token, err := service.UserLoginService(ctx, req)
	println(err)
	if err != nil {
		println("xx")
		return resp, err
	}
	resp.UserId = id
	resp.StatusCode = code
	resp.StatusMsg = &msg
	resp.Token = token
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	resp = &user.UserResponse{
		BaseReponse: &user.BaseResponse{
			StatusCode: 0,
			StatusMsg:  nil,
		},
		User: &user.User{
			Id:            0,
			Name:          "",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		},
	}
	statusCode, statusMsg, u, err := service.UserInfoService(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.BaseReponse.StatusCode = statusCode
	resp.BaseReponse.StatusMsg = &statusMsg
	resp.User = u
	return resp, nil
}
