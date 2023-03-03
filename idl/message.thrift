namespace go message

struct douyinMessageChatRequest {
    1: string token // 用户鉴权token
    2: i64 toUserId // 对方用户id
    3: i64 preMsgTime //上次最新消息的时间（新增字段-apk更新中）
}
struct douyinMessageChatResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
    3: list<message> messageList // 消息列表
}


struct douyinRelationActionRequest {
    1: string token // 用户鉴权token
    2: i64 toUserId // 对方用户id
    3: i32 actionType // 1-发送消息
    4: string content // 消息内容
}
struct douyinRelationActionResponse {
    1: i32 statusCode // 状态码，0-成功，其他值-失败
    2: optional string statusMsg // 返回状态描述
}

struct message {
    1: i64 id // 消息id
    2: i64 toUserId // 该消息接收者的id
    3: i64 fromUserId // 该消息发送者的id
    4: string content // 消息内容
    5: optional string createTime // 消息创建时间
}