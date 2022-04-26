package controller

import (
	"context"
	"path"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gpage"
	"github.com/gogf/gf/v2/util/guid"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cFile struct{}

var File cFile

// Index 控制展示文件上传页面
func (cf *cFile) IndexFiles(ctx context.Context, req *v1.IndexFilesReq) (res *v1.IndexFilesRes, err error) {
	userId := service.Context().Get(ctx).User.Id
	filesMap, err := service.File().GetDirFiles(ctx, userId, req.ParentId, req.Dir, req.Page, req.Size)
	totalSize, _ := service.File().CountDirFiles(ctx, userId, req.ParentId)
	page := g.RequestFromCtx(ctx).GetPage(totalSize, req.Size)
	breadCrumbs := service.View().GetBreadCrumbView(ctx, req.ParentId)
	currentPathId := breadCrumbs[len(breadCrumbs)-1].CurrentPathId
	service.View().Render(ctx, model.View{
		Title: "资源上传",
		Data: g.Map{
			"page":          pageContent(page),
			"filesMap":      filesMap,
			"currentPathId": currentPathId,
		},
		BreadCrumbs: breadCrumbs,
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

func pageContent(page *gpage.Page) string {
	return page.GetContent(3) + "第" + page.SelectBar() + "页"
}

// FileDetail 控制查看文件详情
func (cf *cFile) FileDetail(ctx context.Context, req *v1.FileDetailReq) (res *v1.FileDetailRes, err error) {
	file, err := service.File().GetFileByFileIdAndUserId(ctx, req.FileId, service.Session().GetUser(ctx).Id)
	service.View().Render(ctx, model.View{
		Title: "文件详情",
		Data:  file,
	})
	return
}

// FileRename 控制修改文件名
func (cf *cFile) FileRename(ctx context.Context, req *v1.FileRenameReq) (res *v1.FileRenameRes, err error) {
	if req.NewName == "" {
		return res, gerror.New("文件名不能为空！")
	}
	err = service.File().RenameFile(ctx, service.Session().GetUser(ctx).Id, req.FileId, req.NewName)
	return
}

// FileDownload 控制下载文件
func (cf *cFile) FileDownload(ctx context.Context, req *v1.FileDownloadReq) (res *v1.FileDownloadRes, err error) {
	path, err := service.File().GetFilePathByFileIdAndUserId(ctx, req.FileId, service.Session().GetUser(ctx).Id)
	if err != nil {
		return res, gerror.New("下载失败")
	}
	path = service.File().GetFilesRoot(ctx) + path
	g.RequestFromCtx(ctx).Response.ServeFileDownload(path)
	return
}

// FilePreview 控制预览文件
func (cf *cFile) FilePreview(ctx context.Context, req *v1.FilePreviewReq) (res *v1.FilePreviewRes, err error) {
	path, err := service.File().GetFilePathByFileIdAndUserId(ctx, req.FileId, service.Session().GetUser(ctx).Id)
	if err != nil {
		return res, gerror.New("预览失败")
	}
	path = service.File().GetFilesRoot(ctx) + path
	g.RequestFromCtx(ctx).Response.ServeFile(path)
	return
}

// FileDelete 控制删除文件
func (cf *cFile) FileDelete(ctx context.Context, req *v1.FileDeleteReq) (res *v1.FileDeleteRes, err error) {
	userId := service.Session().GetUser(ctx).Id
	if isFile := service.File().IsFile(ctx, req.FileId); isFile {
		err = service.File().DeleteFileByFileIdAndUserId(ctx, req.FileId, userId)
	} else {
		err = service.Directory().DeleteDirByDirIdAndUserId(ctx, req.FileId, userId)
	}
	return
}

func (cf *cFile) ShowMove(ctx context.Context, req *v1.MoveReq) (res *v1.MoveRes, err error) {
	filesMap, err := service.File().GetDirFiles(ctx, service.Session().GetUser(ctx).Id, req.ParentId, req.Dir, 0, 0)
	service.View().Render(ctx, model.View{
		Title: "移动文件",
		Data: g.Map{
			"filesMap": filesMap,
		},
	})
	return
}

func (cf *cFile) FileMove(ctx context.Context, req *v1.MoveToReq) (res *v1.MoveToRes, err error) {
	dirId := req.DirId
	str := gstr.TrimRight(req.FileIds, ",")
	fileIds := gstr.Split(str, ",")
	userId := service.Session().GetUser(ctx).Id
	dest, _ := service.File().GetFilePathByFileIdAndUserId(ctx, dirId, userId)
	for _, fileId := range fileIds {
		src, _ := service.File().GetFilePathByFileIdAndUserId(ctx, fileId, userId)
		if src == dest || gstr.Contains(dest, src) {
			service.View().Render(ctx, model.View{
				Title: "移动失败",
				Error: "不能将文件移动至自身目录下，请换个目录",
			})
			return res, gerror.New("不能将文件移动至自身目录下，请<a href='/file'>换个目录</a>")
		}
	}
	for _, fileId := range fileIds {
		if service.File().IsFile(ctx, fileId) {
			err = service.File().Move(ctx, userId, fileId, dirId)
		} else {
			err = service.Directory().Move(ctx, userId, fileId, dirId)
		}
		if err != nil {
			service.View().Render(ctx, model.View{
				Title: "移动失败",
				Error: "移动失败，请<a href='/file'>重试</a>",
			})
			return
		}
	}
	service.View().Render(ctx, model.View{
		Title: "移动成功",
		Error: "移动成功，<a href='/file'>点击返回</a>",
	})
	return res, nil
}
