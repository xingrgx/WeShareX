package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xingrgx/WeShareX/internal/model/entity"
)

type IndexAdminReq struct {
	g.Meta `path:"/admin" method:"get"`
}

type IndexAdminRes struct {

}

type EnableReq struct { 
	g.Meta `path:"/admin/enable" method:"post"`
	UserId uint `json:"userId"`
}

type EnableRes struct {

}

type DisableReq struct { 
	g.Meta `path:"/admin/disable" method:"post"`
	UserId uint `json:"userId"`
}

type DisableRes struct {

}

type ProfileReq struct {
	g.Meta `path:"/admin/profile" method:"post"`
	UserId uint `json:"userId"`
}

type ProfileRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type UpdateReq struct {
	g.Meta `path:"/admin/update" method:"post"`
	Input entity.User `json:"input"`
}

type UpdateRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type FileReq struct {
	g.Meta `path:"/admin/file" method:"get"`
	UserId uint `json:"userId"`
	ParentId string `json:"dirId" d:"root"`
	Dir      int    `json:"dir" d:"0" summary:"Dir=0:展示文件和文件夹；dir=1:仅展示文件夹"`
	PaginationReq
}

type FileRes struct {
	PaginationRes
}

type AdminFileDeleteReq struct {
	g.Meta `path:"/admin/file/delete" method:"post" summary:"预览文件" tags:"文件"`
	FileId string `json:"fileId"`
	UserId uint `json:"userId"`
}

type AdminFileDeleteRes struct {

}

type AdminShareReq struct {
	g.Meta `path:"/admin/share" method:"get"`
	UserId uint `json:"userId"`
}

type AdminShareRes struct {

}