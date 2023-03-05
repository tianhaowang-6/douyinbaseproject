namespace go message

struct MessageChatRequest {
    1: string token (go.tag='json:"token"')// 用户鉴权token
    2: i64 toUserId (go.tag='json:"to_user_id"')// 对方用户id
    3: i64 preMsgTime (go.tag='json:"pre_msg_time"')//上次最新消息的时间（新增字段-apk更新中）
}
struct MessageChatResponse {
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg (go.tag='json:"status_msg"')// 返回状态描述
    3: list<message> messageList (go.tag='json:"message_list"')// 消息列表
}


struct RelationActionMessageRequest {
    1: string token (go.tag='json:"token"')// 用户鉴权token
    2: i64 toUserId (go.tag='json:"to_user_id"')// 对方用户id
    3: i32 actionType(go.tag='json:"action_type"') // 1-发送消息
    4: string content (go.tag='json:"content"')// 消息内容
}
struct RelationActionMessageResponse {
    1: i32 statusCode (go.tag='json:"status_code"')// 状态码，0-成功，其他值-失败
    2: optional string statusMsg(go.tag='json:"status_msg"') // 返回状态描述
}

struct message {
    1: i64 id (go.tag='json:"id"')// 消息id
    2: i64 toUserId (go.tag='json:"to_user_id"')// 该消息接收者的id
    3: i64 fromUserId(go.tag='json:"from_user_id"')// 该消息发送者的id
    4: string content(go.tag='json:"content"')// 消息内容
    5: optional string createTime(go.tag='json:"create_time"') // 消息创建时间
}

service MessageService{
    MessageChatResponse  MessageChat(MessageChatRequest req)
    RelationActionMessageResponse RelationActionMessage(RelationActionMessageRequest req)
}