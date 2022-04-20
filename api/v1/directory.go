package v1

import "github.com/gogf/gf/v2/frame/g"

type DirUploadReq struct {
	g.Meta `path:"/dir/create" method:"post" summary:"上传文件夹" tags:"文件夹"`
	Name string `json:"name"`
	ParentId string `json:"parentId"`
}

type DirUploadRes struct {

}
