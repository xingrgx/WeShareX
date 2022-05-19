package model

// ChatMsg 公共聊天室信息结构
type ChatMsg struct {
	Type string      `json:"type" v:"required"`
	Data interface{} `json:"data" v:"required"`
	From string      `json:"name" v:""`
}

// FriendProfile 通讯录展示的好友信息
type FriendProfile struct {
	Id       uint
	Passport string
	Nickname string
	Gender   int
	Status   int
}
