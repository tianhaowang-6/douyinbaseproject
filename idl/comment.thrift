namespace go comment

struct CommentActionRequest {
    1: string token // 用户鉴权token
    2: i64 videoId // 视频id
    3: i32 actionType // 1-发布评论，2-删除评论
    4: optional string commentText // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 commentId // 要删除的评论id，在action_type=2的时候使用
}
struct CommentActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: optional comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
}
struct comment {
    1: i64 id // 视频评论id
    2: user user // 评论用户信息
    3: string content // 评论内容
    4: string createDate // 评论发布日期，格式 mm-dd
}

struct CommentListRequest {
    1: string token // 用户鉴权token
    2: i64 videoId // 视频id
}
struct CommentListResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<comment> commentList // 评论列表
}
struct user {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: optional i64 followCount // 关注总数
    4: optional i64 followerCount // 粉丝总数
    5: bool isFollow // true-已关注，false-未关注
}
service CommentService{
    CommentActionResponse CommentAction (1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
}
