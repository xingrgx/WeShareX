package controller

import (
	"context"
	"io"
	"path"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
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
	filesMap, err := service.File().GetDirFiles(ctx, userId, req.ParentId, req.Page, req.Size)
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

func (cd *cFile) UploadBigFile(ctx context.Context, req *v1.UploadBigFileReq) (res *v1.UploadBigFileRes, err error) {
	r := g.RequestFromCtx(ctx)
	g.Client().SetTimeout(6000 * time.Second)
	r.Server.SetClientMaxBodySize(1024 * 1024 * 1024 * 10)
	file := r.GetUploadFile("file")
	name, err := file.Save(req.Path)
	if err == nil {
		g.Dump(name)
	}
	return
}

// UploadFiles 控制上传文件
func (cf *cFile) UploadFile(ctx context.Context, req *v1.UploadReq) (*v1.UploadRes, error) {
	var res *v1.UploadRes
	g.Client().SetTimeout(5 * time.Second)
	if err := g.RequestFromCtx(ctx).ParseMultipartForm(128<<20); err != nil {
		return res, err
	}
	if srcFile, fileHeader, err := g.RequestFromCtx(ctx).FormFile("file"); err == nil {
		defer func() { _ = srcFile.Close() } ()
		fileType := path.Ext(fileHeader.Filename)
		fileInput := model.FileUploadInput{
			Id:       guid.S(),
			UserId:   service.Session().GetUser(ctx).Id,
			Name:     fileHeader.Filename,
			ParentId: req.ParentId,
			Dir:      req.Dir,
			Type:     fileType,
			Size:     fileHeader.Size,
		}
		relPath, _ := service.File().GetFileRelativePath(ctx, fileInput.UserId, fileInput.ParentId, fileInput.Name)
		fileInput.Path = relPath
		if err = service.File().UploadFile(ctx, fileInput); err != nil {
			return res, err
		}
		absPath := service.File().GetFileAbsoluteParentPath(ctx, fileInput.UserId, fileInput.ParentId) + fileInput.Name
		g.Dump(absPath)
		dstFile, err := gfile.Create(absPath)
		if err != nil {
			return res, err
		}
		defer func() { dstFile.Close() }()
		if _, err = io.Copy(dstFile, srcFile); err != nil {
			return res, err
		}
	} else {
		return res, err
	}

	//err = service.File().Save(file, absPPath)
	return res, nil
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
