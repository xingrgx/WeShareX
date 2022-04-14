package controller

import (
	"context"
	"path"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cFile struct{}

var File cFile

// Index 控制展示文件上传页面
func (cf *cFile) Index(ctx context.Context, req *v1.IndexUploadReq) (res *v1.IndexUploadRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "资源上传",
	})
	return
}

// UploadFiles 控制上传文件
func (cf *cFile) UploadFile(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	file := g.RequestFromCtx(ctx).GetUploadFile("file")
	fileType := path.Ext(file.Filename)
	fileInput := model.FileUploadInput{
		Id:       guid.S(),
		UserId:   service.Session().GetUser(ctx).Id,
		Name:     file.Filename,
		ParentId: req.ParentId,
		Dir:      req.Dir,
		Type:     fileType,
		Size:     file.Size,
	}
	relPath, _ := service.File().GetFileRelativePath(ctx, fileInput.UserId, fileInput.ParentId, fileInput.Name)
	fileInput.Path = relPath
	if err = service.File().UploadFile(ctx, fileInput); err != nil {
		return res, err
	}
	absPPath := service.File().GetFileAbsoluteParentPath(ctx, fileInput.UserId, fileInput.ParentId)
	err = service.File().Save(file, absPPath)
	return
}
