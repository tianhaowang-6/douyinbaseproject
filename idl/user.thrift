namespace go user


struct baseResponse{
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}


struct douyinUserRegisterRequest {
    1: string username // 注册用户名，最长32个字符
    2: string password // 密码，最长32个字符
}
struct douyinUserRegisterResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: i64 userId // 用户id
    4: string token // 用户鉴权token
}

struct douyinUserLoginRequest {
    1: string username // 登录用户名
    2: string password // 登录密码
}
struct douyinUserLoginResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: i64 userId // 用户id
    4: string token // 用户鉴权token
}

struct douyinUserRequest {
    1: i64 userId // 用户id
    2: string token // 用户鉴权token
}
struct douyinUserResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: UserInfo user // 用户信息
}



struct UserInfo {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: optional i64 followCount // 关注总数
    4: optional i64 followerCount // 粉丝总数
    5: bool isFollow // true-已关注，false-未关注
}

service UserService{
    douyinUserRegisterResponse CreateUser(1:douyinUserRegisterRequest req)
    douyinUserLoginResponse CheckUser(1:douyinUserLoginRequest req)
    douyinUserResponse GetUserInfo(douyinUserRequest req)

}
