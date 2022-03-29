package v1

import "github.com/gogf/gf/v2/frame/g"

// OpenAPIv3 自动生成 api 接口文档

// LoginIndexReq 首页跳转到登录页面的请求
type LoginIndexReq struct {
	g.Meta `path:"/login" method:"get" summary:"展示登录页面" tags:"登录"`
}

// LoginIndexRes 首页跳转到登录页面的响应
type LoginIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// LoginReq 登录页面的登录请求
type LoginReq struct {
	g.Meta          `path:"/login" method:"post" sm:"执行登录请求" tags:"登录"`
	Passport string `json:"passport" v:"required#请输入账号" sm:"账号"`
	Password string `json:"password" v:"required#请输入密码" sm:"密码（明文）"`
}

// LoginRes 登录页面的登录响应
type LoginRes struct {
	Referer string `json:"referer" dc:"引导客户端跳转地址"`
}