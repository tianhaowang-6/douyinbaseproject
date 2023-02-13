namespace go feed

struct douyinFeedRequest {
    1: optional i64 latestTime // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token // 可选参数，登录用户设置
}

struct douyinFeedResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<video> videoList // 视频列表
    4: optional i64 nextTime // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct video {
    1: i64 id // 视频唯一标识
    2: user author // 视频作者信息
    3: string playUrl // 视频播放地址
    4: string coverUrl // 视频封面地址
    5: i64 favoriteCount // 视频的点赞总数
    6: i64 commentCount // 视频的评论总数
    7: bool isFavorite // true-已点赞，false-未点赞
    8: string title // 视频标题
}
struct user {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: optional i64 followCount // 关注总数
    4: optional i64 followerCount // 粉丝总数
    5: bool isFollow // true-已关注，false-未关注
}
struct douyinPublishActionRequest {
    1: string token // 用户鉴权token
    2: byte data // 视频数据
    3: string title // 视频标题
}

struct douyinPublishActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}
