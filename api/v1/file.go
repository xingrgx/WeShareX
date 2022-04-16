package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// IndexFilesReq 展示上传文件页面的请求
type IndexFilesReq struct {
	g.Meta `path:"/file" method:"get" summary:"展示上传页面" tags:"文件"`
	PaginationReq
}

// IndexFilesRes 展示上传文件页面的响应
type IndexFilesRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	PaginationRes
}

// UploadReq 上传文件的请求
type UploadReq struct {
	g.Meta   `path:"/file/upload" method:"post" summary:"上传文件" tags:"文件"`
	ParentId string `json:"parentId"`
	Dir      int    `json:"dir"`
}

type UploadRes struct {
}

// FileDetailReq 查看文件详情的请求接口
type FileDetailReq struct {
	g.Meta   `path:"/file/detail" method:"get" summary:"查看文件详情" tags:"文件"`
	FileId string `json:"fileId"`
}

type FileDetailRes struct {
	
}

// FileRenameReq 修改文件名的请求接口
type FileRenameReq struct {
	g.Meta `path:"/file/rename" method:"post" summary:"修改文件名" tags:"文件"`
	FileId string `json:"fileId"`
	NewName string `json:"newName"`
}

type FileRenameRes struct {

}

// FileDownloadReq 下载文件的请求接口
type FileDownloadReq struct {
	g.Meta `path:"/file/download" method:"get" summary:"下载文件" tags:"文件"`
	FileId string `json:"fileId"` 
}

type FileDownloadRes struct {

}

// FilePreviewReq 预览文件的请求接口
type FilePreviewReq struct {
	g.Meta `path:"/file/preview" method:"get" summary:"预览文件" tags:"文件"`
	FileId string `json:"fileId"` 
}

type FilePreviewRes struct {

}

// FileDeleteReq 删除文件的请求接口
type FileDeleteReq struct {
	g.Meta `path:"/file/delete" method:"post" summary:"预览文件" tags:"文件"`
	FileId string `json:"fileId"` 
}

type FileDeleteRes struct {
	
}