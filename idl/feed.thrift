namespace go feed

struct  FeedRequest {
    1: optional i64 latestTime (go.tag='json:""')// 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token(go.tag='json:"token"') // 可选参数，登录用户设置
}

struct  FeedResponse {
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg(go.tag='json:"status_msg"') // 返回状态描述
    3: list<video> videoList(go.tag='json:"video_list"') // 视频列表
    4: optional i64 nextTime(go.tag='json:"next_time"') // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct video {
    1: i64 id (go.tag='json:"id"')// 视频唯一标识
    2: user author (go.tag='json:"author"')// 视频作者信息
    3: string playUrl(go.tag='json:"play_url"') // 视频播放地址
    4: string coverUrl (go.tag='json:"cover_url"')// 视频封面地址
    5: i64 favoriteCount(go.tag='json:"favorite_count"')// 视频的点赞总数
    6: i64 commentCount (go.tag='json:"comment_count"')// 视频的评论总数
    7: bool isFavorite(go.tag='json:"is_favorite"')// true-已点赞，false-未点赞
    8: string title(go.tag='json:"title"')// 视频标题
}
struct user {
    1: i64 id (go.tag='json:"id"')// 用户id
    2: string name(go.tag='json:"name"') // 用户名称
    3: optional i64 followCount(go.tag='json:"follow_count"')// 关注总数
    4: optional i64 followerCount(go.tag='json:"follower_count"')// 粉丝总数
    5: bool isFollow (go.tag='json:"is_follow"')// true-已关注，false-未关注
}
struct  PublishActionRequest {
    1: string token (go.tag='json:"token"')// 用户鉴权token
    2: string  filePath(go.tag='json:"file_path"')// 文件路径
    3: string  coverPath(go.tag='json:"cover_path"')// 文件路径
    4: string title (go.tag='json:"title"')// 视频标题
}
struct  PublishActionResponse {
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg(go.tag='json:"status_msg"') // 返回状态描述
}

struct PublishListRequest {
    1: i64 userId (go.tag='json:"user_id"')// 用户id
    2: string token(go.tag='json:"token"')// 用户鉴权token
}
struct PublishListResponse {
    1: i32 statusCode(go.tag='json:"status_code"') // 状态码，0-成功，其他值-失败
    2: optional string statusMsg(go.tag='json:"status_msg"') // 返回状态描述
    3: list<video> videoList(go.tag='json:"video_list"')// 用户发布的视频列表
}
//
struct FavoriteActionRequest {
    1: string token(go.tag='json:"token"') // 用户鉴权token
    2: i64 videoId(go.tag='json:"video_id"') // 视频id
    3: i32 actionType(go.tag='json:"action_type"') // 1-点赞，2-取消点赞
}
struct FavoriteActionResponse {
    1: i32 statusCode(go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg(go.tag='json:"status_msg"')// 返回状态描述
}
//
struct FavoriteListRequest {
    1: i64 userId(go.tag='json:"user_id"')// 用户id
    2: string token(go.tag='json:"token"') // 用户鉴权token
}
struct FavoriteListResponse {
    1: i32 statusCode(go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg(go.tag='json:"status_msg"') // 返回状态描述
    3: list<video> videoList(go.tag='json:"video_list"') // 用户点赞视频列表
}

service FeedService{
    // 获取视频流
    FeedResponse GetFeed(1 : FeedRequest req)
    // 上传视频流
    PublishActionResponse PublishAction(1:PublishActionRequest req)
    // 发布列表
    PublishListResponse PublishList(1:PublishListRequest req)
    // 赞操作
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    // 喜欢列表
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
}