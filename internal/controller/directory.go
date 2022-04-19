package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
	"github.com/xingrgx/WeShareX/utility/response"
)

type cDirectory struct{}

var Directory cDirectory

// DirUpload 控制上传文件夹
func (cd *cDirectory) DirUpload(ctx context.Context, req *v1.DirUploadReq) (res *v1.DirUploadRes, err error) {
	userId := service.Session().GetUser(ctx).Id
	dirId := guid.S()
	dir := model.FileUploadInput{
		Id:       dirId,
		UserId:   userId,
		Name:     req.Name,
		ParentId: req.ParentId,
		Dir:      1,
	}
	dir.Path, _ = service.File().GetFileRelativePath(ctx, userId, req.ParentId, req.Name)
	// 当前文件夹不存在才上传
	if e := service.File().CheckFileNameExist(ctx, dir.Name, dir.UserId, dir.Path); e == nil {
		service.File().UploadFile(ctx, dir)
	}
	// 查询当前文件夹信息以返回
	dirInfo, _ := service.File().GetFileByFileNamePathAndUserId(ctx, dir.Name, dir.Path, userId)
	response.Json(g.RequestFromCtx(ctx), 0, "", gconv.Map(dirInfo))
	return
}
