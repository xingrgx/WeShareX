package controller

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
	if req.Name == "" {
		return res, gerror.New("文件夹名不能为空")
	}
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

// 控制下载多个文件夹或文件
func (cd *cDirectory) DirsAndFilesZip(ctx context.Context, req *v1.ZipReq) (res *v1.ZipRes, err error) {
	g.Dump(req.FileIds)
	str := gstr.TrimRight(req.FileIds, ",")
	fileIds := gstr.Split(str, ",")
	g.Dump(fileIds)
	userId := service.Session().GetUser(ctx).Id
	paths := []string{}
	root := service.File().GetFilesRoot(ctx)
	zipPath := root + "/files.zip"
	for _, fileId := range fileIds {
		path, _ := service.File().GetFilePathByFileIdAndUserId(ctx, fileId, userId)
		paths = append(paths, root+path)
	}
	service.Directory().Zip(paths, zipPath)
	g.RequestFromCtx(ctx).Response.ServeFileDownload(zipPath)
	defer os.Remove(zipPath)
	return
}
