namespace go user


struct BaseResponse{
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg (go.tag='json:"status_msg"') // 返回状态描述
}
struct UserRegisterRequest {
    1: string username  (go.tag='json:"username"')// 注册用户名，最长32个字符
    2: string password  (go.tag='json:"password"')// 密码，最长32个字符
}
struct UserRegisterResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg  (go.tag='json:"status_msg"')// 返回状态描述
    3: i64 userId (go.tag='json:"user_id"') // 用户id
    4: string token  (go.tag='json:"token"')// 用户鉴权token
}

struct UserLoginRequest {
    1: string username  (go.tag='json:"username"')// 登录用户名
    2: string password  (go.tag='json:"password"')// 登录密码
}
struct UserLoginResponse {
    1: i32 statusCode (go.tag='json:"status_code"') // 状态码，0-成功，其他值-失败
    2: optional string statusMsg (go.tag='json:"status_msg"')// 返回状态描述
    3: i64 userId (go.tag='json:"user_id"') // 用户id
    4: string token (go.tag='json:"token"')// 用户鉴权token
}

struct UserRequest {
    1: i64 userId (go.tag='json:"user_id"') // 用户id
    2: string token  (go.tag='json:"token"')// 用户鉴权token
}
struct UserResponse {
    1: BaseResponse baseReponse (go.tag='json:"base_reponse"')
    3: User user  (go.tag='json:"user"')// 用户信息
}

struct User {
    1: i64 id  (go.tag='json:"id"')// 用户id
    2: string name (go.tag='json:"name"')// 用户名称
    3: optional i64 followCount (go.tag='json:"follow_count"') // 关注总数
    4: optional i64 followerCount (go.tag='json:"follower_count"')// 粉丝总数
    5: bool isFollow (go.tag='json:"is_follow"') // true-已关注，false-未关注
}

service UserService{
    // 注册
    UserRegisterResponse UserRegister(1:UserRegisterRequest req)
    // 登录
    UserLoginResponse UserLogin(1:UserLoginRequest req)
    UserResponse UserInfo(UserRequest req)
}
