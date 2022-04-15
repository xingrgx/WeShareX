package service

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xingrgx/WeShareX/internal/consts"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/model/entity"
	"github.com/xingrgx/WeShareX/internal/service/internal/dao"
)

type sFile struct{}

// File 获取文件服务
func File() *sFile {
	return &sFile{}
}

// UploadFiles 上传文件
func (sf *sFile) UploadFile(ctx context.Context, fileInput model.FileUploadInput) error {
	return dao.File.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var file *entity.File
		if err := gconv.Struct(fileInput, &file); err != nil {
			return err
		}
		// 检测文件是否已存在
		if err := sf.CheckFileNameExist(ctx, fileInput.Name, fileInput.UserId, fileInput.Path); err != nil {
			return err
		}
		_, err := dao.File.Ctx(ctx).Data(file).OmitEmpty().Save()
		return err
	})
}

// CheckFileNameExist 检查文件名是否存在，检测规则：Filename+UserID+Path
func (sf *sFile) CheckFileNameExist(ctx context.Context, filename string, userId uint, path string) error {
	n, err := dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().Name:   filename,
		dao.File.Columns().UserId: userId,
		dao.File.Columns().Path:   path,
	}).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf("文件%s已存在", filename)
	}
	return nil
}

// Save 保存文件至服务器磁盘
func (sf *sFile) Save(file *ghttp.UploadFile, path string) error {
	_, err := file.Save(path)
	return err
}

// GetFilesRoot 获取当前用户的根路径，默认为 files/{userId}/root
func (sf *sFile) GetFilesRoot(ctx context.Context) (rootPath string) {
	filePath := g.Cfg().MustGet(ctx, "server.filePath").String()
	userId := Session().GetUser(ctx).Id
	rootPath = fmt.Sprintf("%s/%d/%s", filePath, userId, consts.FILES_ROOT)
	return
}

// GetFileRelativeParentPath 根据当前文件的UserId,ParentId获取其父级相对路径
// 该路径用于存储到file表中的path字段
// e.g. root -> /
// e.g. /hello
func (sf *sFile) GetFileRelativeParentPath(ctx context.Context, userId uint, parentId string) (path string, err error) {
	if parentId == "root" {
		return "/", nil
	}
	var file *entity.File
	err = dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().Id:     parentId,
		dao.File.Columns().UserId: userId,
	}).Scan(&file)
	if err != nil {
		return "", err
	}
	arr := gstr.Split(file.Path, "/")
	arr = arr[:len(arr)-1]
	prtPath := gstr.Join(arr, "/")
	return prtPath, nil
}

// GetFileRelativePath 获取当前文件的相对路径
func (sf *sFile) GetFileRelativePath(ctx context.Context, userId uint, parentId, name string) (path string, err error) {
	if parentId == "root" {
		return "/" + name, nil
	}
	prtPath, err := sf.GetFileRelativeParentPath(ctx, userId, parentId)
	if err != nil {
		return "", err
	}
	return prtPath + "/" + name, nil
}

// GetUserFileAbsolutePath 获取当前文件的父级绝对路径
func (sf *sFile) GetFileAbsoluteParentPath(ctx context.Context, userId uint, parentId string) (path string) {
	prtPath, _ := sf.GetFileRelativeParentPath(ctx, userId, parentId)
	return sf.GetFilesRoot(ctx) + "/" + prtPath
}

// GetRootFiles 获取root目录下的所有文件
func (sf *sFile) GetRootFiles(ctx context.Context, userId uint) (filesMap []g.Map, err error) {
	var filesArr []entity.File
	dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().ParentId: "root",
		dao.File.Columns().UserId:   userId,
	}).Scan(&filesArr)
	for _, file := range filesArr {
		filesMap = append(filesMap, gconv.Map(file))
	}
	return
}
