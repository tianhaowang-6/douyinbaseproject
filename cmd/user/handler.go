package main

import (
	"context"
	"douyin-pro/cmd/user/service"
	user "douyin-pro/kitex_gen/user"
	"douyin-pro/pkg/errno"
	"fmt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserRegisterResponse)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
		return resp, nil
	}
	service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		return
	}
	resp.StatusCode = errno.SuccessCode
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	println("checkuser", req.Username, req.Password)
	resp = new(user.DouyinUserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
		return resp, nil
	}
	id, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		fmt.Printf("查询失败完成进行返回%#v\n", err)
		resp.StatusCode = errno.AuthorizationFailedErrCode
		resp.StatusMsg = &errno.AuthorizationFailedErr.ErrMsg
		return resp, nil
	}
	resp.UserId = id
	//resp.Token=
	resp.StatusCode = errno.SuccessCode
	fmt.Printf("查询完成进行返回%#v\n", resp)
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	userId := req.UserId
	resp = new(user.DouyinUserResponse)
	if userId == 0 {
		resp.StatusCode = errno.ErrorCode
		resp.StatusMsg = &errno.Error.ErrMsg
		return resp, nil
	}
	userInfo, err := service.NewQueryUserService(ctx).QueryUserInfo(req)
	if err != nil {
		fmt.Printf("查询失败完成进行返回%#v\n", err)
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = &errno.ParamErr.ErrMsg
		return resp, nil
	}
	resp.User = userInfo
	return resp, nil
}
