namespace go comment

struct CommentActionRequest {
    1: string token(go.tag='json:"token"') // 用户鉴权token
    2: i64 videoId (go.tag='json:"video_id"')// 视频id
    3: i32 actionType (go.tag='json:"action_type"')// 1-发布评论，2-删除评论
    4: optional string commentText(go.tag='json:"comment_text"') // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 commentId(go.tag='json:"comment_id"') // 要删除的评论id，在action_type=2的时候使用
}
struct CommentActionResponse {
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg (go.tag='json:"status_msg"')// 返回状态描述
    3: optional Comment comment (go.tag='json:"comment"')// 评论成功返回评论内容，不需要重新拉取整个列表
}
struct Comment {
    1: i64 id (go.tag='json:"id"') // 视频评论id
    2: user user(go.tag='json:"user"')  // 评论用户信息
    3: string content (go.tag='json:"content"') // 评论内容
    4: string createDate (go.tag='json:"create_date"') // 评论发布日期，格式 mm-dd
    5: i64 videoId(go.tag='json:"video_id"')
}

struct CommentListRequest {
    1: string token(go.tag='json:"token"') // 用户鉴权token
    2: i64 videoId (go.tag='json:"video_id"')// 视频id
}

struct CommentListResponse {
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg (go.tag='json:"status_msg"')// 返回状态描述
    3: list<Comment> commentList (go.tag='json:"comment_list"')// 评论列表
}
struct user {
    1: i64 id(go.tag='json:"id"') // 用户id
    2: string name(go.tag='json:"name"') // 用户名称
    3: optional i64 followCount (go.tag='json:"follow_count"')// 关注总数
    4: optional i64 followerCount (go.tag='json:"follower_count"')// 粉丝总数
    5: bool isFollow(go.tag='json:"is_follow"') // true-已关注，false-未关注
}
service CommentService{
    CommentActionResponse CommentAction (1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
}
