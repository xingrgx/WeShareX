package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// IndexFilesReq 展示上传文件页面的请求
type IndexFilesReq struct {
	g.Meta `path:"/file" method:"get" summary:"展示上传页面" tags:"文件"`
}

// IndexFilesRes 展示上传文件页面的响应
type IndexFilesRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// UploadReq 上传文件的请求
type UploadReq struct {
	g.Meta   `path:"/file/upload" method:"post" summary:"上传文件" tags:"文件"`
	ParentId string `json:"parentId"`
	Dir      int    `json:"dir"`
}

type UploadRes struct {
}
