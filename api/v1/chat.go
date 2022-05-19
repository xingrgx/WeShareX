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
	Id uint `json:"id"`
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

type ListFriendsReq struct {
	g.Meta `path:"/chat/friends" method:"get" tags:"通讯" summary:"查询通讯录"`
}

type ListFriendsRes struct {
	
}

type AddFriendReq struct {
	g.Meta `path:"/chat/add" method:"post" tags:"通讯" summary:"添加好友"`
	Id string `json:"id"`
}

type AddFriendRes struct {
	
}

type SearchFriendReq struct {
	g.Meta `path:"/chat/search" method:"get" tags:"通讯" summary:"查找好友"`
	Msg string `json:"msg"`
}

type SearchFriendRes struct {
	
}

type AgreeReq struct {
	g.Meta `path:"/chat/agree" method:"get" tags:"通讯" summary:"同意请求"`
	Id string `json:"id"`
}

type AgreeRes struct {
	
}