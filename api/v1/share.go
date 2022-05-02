package v1

import "github.com/gogf/gf/v2/frame/g"

type ShareToCreateReq struct {
	g.Meta  `path:"/share/create" method:"get" summary:"展示分享页面" tags:"分享"`
	FileIds string `json:"fileIds"`
}

type ShareToCreateRes struct {
}

type ShareCreateReq struct {
	g.Meta      `path:"/share/create" method:"post" summary:"展示分享页面" tags:"分享"`
	FileIds     string `json:"fileIds"`
	NeverExpire string `json:"neverExpire"`
	ExpireTime  string `json:"expireTime"`
}

type ShareCreateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type ShareListReq struct {
	g.Meta `path:"/share" method:"get" summary:"获取所有分享" tags:"分享"`
}

type ShareListRes struct {
}

type ShareLinkReq struct {
	g.Meta `path:"/share/link" method:"get" summary:"获取分享链接" tags:"分享"`
	Id     string `json:"id" sm:"分享的ID"`
}

type ShareLinkRes struct {
}

