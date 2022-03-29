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