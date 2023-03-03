namespace go relationship

struct RelationActionRequest {
1: string token // 用户鉴权token
2: i64 toUserId // 对方用户id
3: i32 actionType // 1-关注，2-取消关注
}
struct RelationActionResponse {
1: i32 statusCode // 状态码，0-成功，其他值-失败
2: optional string statusMsg // 返回状态描述
}

struct RelationFollowListRequest {
1: i64 userId // 用户id
2: string token // 用户鉴权token
}
struct RelationFollowListResponse {
1: i32 statusCode // 状态码，0-成功，其他值-失败
2: optional string statusMsg // 返回状态描述
3: list<user> userList // 用户信息列表
}
struct user {
1: i64 id // 用户id
2: string name // 用户名称
3: optional i64 followCount // 关注总数
4: optional i64 followerCount // 粉丝总数
5: bool isFollow // true-已关注，false-未关注
}

struct RelationFollowerListRequest {
1: i64 userId // 用户id
2: string token // 用户鉴权token
}
struct RelationFollowerListResponse {
1: i32 statusCode // 状态码，0-成功，其他值-失败
2: optional string statusMsg // 返回状态描述
3: list<user> userList // 用户列表
}

struct RelationFriendListRequest {
    1: i64 userId // 用户id
    2: string token // 用户鉴权token
}
struct RelationFriendListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<user> userList // 用户列表
}

service RelationShipService{
    RelationActionResponse RelationAction(1:RelationActionRequest req)
    RelationFollowListResponse RelationFollowList(1:RelationFollowListRequest req)
    RelationFollowerListResponse RelationFollowerList (1:RelationFollowerListResponse req)
    RelationFriendListResponse RelationFriendList (1:RelationFriendListRequest req)
}



