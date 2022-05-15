package model

// ChatMsg 公共聊天室信息结构
type ChatMsg struct {
	Type string      `json:"type" v:"required"`
	Data interface{} `json:"data" v:"required"`
	From string      `json:"name" v:""`
}