package v1

import "github.com/gogf/gf/v2/frame/g"

type IndexChatFriendsReq struct {
	g.Meta `path:"/chat/friends" method:"get" summary:"展示通讯录页面" tags:"通讯"`
}

type IndexChatFriendsRes struct {

}

type IndexChatRoomReq struct {
	g.Meta `path:"/chat/room" method:"get" summary:"展示公共聊天室页面" tags:"通讯"`
}

type IndexChatRoomRes struct {

}

type ChatIndexReq struct {
	g.Meta `path:"/chat" method:"get"  tags:"ChatService" summary:"Home page"`
}
type ChatIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ChatNameReq struct {
	g.Meta `path:"/chat/name" method:"post"  tags:"ChatService" summary:"Name page"`
	Name   string `v:"required|max-length:21#Why not an awesome name"`
}
type ChatNameRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ChatWebsocketReq struct {
	g.Meta `path:"/chat/websocket" method:"get"  tags:"ChatService" summary:"Send message"`
}
type ChatWebsocketRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>" dc:"It redirects to homepage if success"`
}
