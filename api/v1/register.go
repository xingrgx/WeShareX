package v1

import "github.com/gogf/gf/v2/frame/g"

// RegisterIndexReq 首页跳转至注册页面的请求
type RegisterIndexReq struct {
	g.Meta `path:"/register" method:"get" summary:"展示注册页面" tags:"注册"`
}

// RegisterIndexRes 首页跳转至注册页面的响应
type RegisterIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// RegisterReq 注册请求
type RegisterReq struct {
	g.Meta          `path:"/register" method:"post" summary:"执行注册请求" tags:"注册"`
	Passport string `json:"passport" v:"required#请输入账号"    dc:"账号"`
	Nickname string `json:"nickname" v:"required#请输入昵称"    dc:"昵称"`
	Password string `json:"password" v:"required#请输入密码"    dc:"密码（明文）"`
	Email string    `json:"email"    v:"required#请输入邮箱地址" dc:"邮箱"`
}

// RegisterRes 注册响应
type RegisterRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}