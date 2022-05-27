package v1

import "github.com/gogf/gf/v2/frame/g"

// IndexReq 展示首页的请求接口
type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
}

// IndexRes 展示首页的响应接口
type IndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type TestReq struct{
	g.Meta `path:"/test" method:"get"`
}

type TestRes struct{}