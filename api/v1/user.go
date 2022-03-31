package v1

import "github.com/gogf/gf/v2/frame/g"

// UserLogoutReq 用户注销请求接口
type UserLogoutReq struct {
	g.Meta `path:"/user/logout" method:"get" summary:"用户注销请求接口" tags:"个人"`
}

// UserLogoutRes 用户注销响应
type UserLogoutRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// UserHomeReq 用户个人主页请求
type UserHomeReq struct {
	g.Meta `path:"/user/{UserId}" method:"get" summary:"显示用户个人主页" tags:"个人"`
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}

// UserHomeRes 用户个人主页响应
type UserHomeRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}