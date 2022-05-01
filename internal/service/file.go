package service

import (
	"context"
	"fmt"
	"os"

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
		_, err := dao.File.Ctx(ctx).Data(file).Save()
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
	if file.Dir == 1 {
		return file.Path, nil
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

// GetDirFiles 获取指定文件夹下的所有文件
func (sf *sFile) GetDirFiles(ctx context.Context, userId uint, parentId string, dir, page, size int) (filesMap []g.Map, err error) {
	var filesArr []entity.File
	if dir == 0 {
		dao.File.Ctx(ctx).Where(g.Map{
			dao.File.Columns().UserId:   userId,
			dao.File.Columns().ParentId: parentId,
		}).Page(page, size).OrderDesc(dao.File.Columns().UpdateAt).Scan(&filesArr)
	}
	if dir == 1 {
		dao.File.Ctx(ctx).Where(g.Map{
			dao.File.Columns().UserId:   userId,
			dao.File.Columns().ParentId: parentId,
			dao.File.Columns().Dir:      1,
		}).Page(page, size).OrderDesc(dao.File.Columns().UpdateAt).Scan(&filesArr)
	}
	for _, file := range filesArr {
		filesMap = append(filesMap, gconv.Map(file))
	}
	return
}

// CountDirFiles 获取指定目录下的文件数量
func (sf *sFile) CountDirFiles(ctx context.Context, userId uint, parentId string) (totalSize int, err error) {
	totalSize, err = dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().ParentId: parentId,
		dao.File.Columns().UserId:   userId,
	}).Count()
	return
}

// GetFileByFileIdAndUserId 根据用户ID和文件ID获取文件
func (sf *sFile) GetFileByFileIdAndUserId(ctx context.Context, fid string, uid uint) (file *entity.File, err error) {
	err = dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().Id:     fid,
		dao.File.Columns().UserId: uid,
	}).Scan(&file)
	return
}

// RenameFile 修改文件名
func (sf *sFile) RenameFile(ctx context.Context, userId uint, fileId, newName string) (err error) {
	_, err = dao.File.Ctx(ctx).OmitEmpty().Data(dao.File.Columns().Name, newName).
		Where(g.Map{
			dao.File.Columns().Id:     fileId,
			dao.File.Columns().UserId: userId,
		}).Update()
	if err != nil {
		return gerror.New("修改文件名失败！")
	}
	return
}

// GetFilePathByFileIdAndUserId 根据fileId和userId获取path
func (sf *sFile) GetFilePathByFileIdAndUserId(ctx context.Context, fileId string, userId uint) (path string, err error) {
	val, err := dao.File.Ctx(ctx).Fields(dao.File.Columns().Path).Where(g.Map{
		dao.File.Columns().Id:     fileId,
		dao.File.Columns().UserId: userId,
	}).Value()
	if err != nil {
		return "", err
	}
	return val.String(), nil
}

// IsFile 根据ID判断是文件（true）还是文件夹（false）
func (sf *sFile) IsFile(ctx context.Context, id string) (is bool) {
	dir, _ := dao.File.Ctx(ctx).Fields(dao.File.Columns().Dir).Where(g.Map{
		dao.File.Columns().Id: id,
	}).Value()
	return dir.Int() == 0
}

// DeleteFileByFileIdAndUserId 根据filedId和userId删除文件
func (sf *sFile) DeleteFileByFileIdAndUserId(ctx context.Context, fileId string, userId uint) (err error) {
	deleteFileFromDisk(ctx, fileId, userId)
	_, err = dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().Id:     fileId,
		dao.File.Columns().UserId: userId,
	}).Delete()
	if err != nil {
		return gerror.New("删除失败，请重试！")
	}
	return
}

func deleteFileFromDisk(ctx context.Context, fileId string, userId uint) (err error) {
	file, _ := File().GetFileByFileIdAndUserId(ctx, fileId, userId)
	root := File().GetFilesRoot(ctx)
	filePath := "./" + root + file.Path
	os.Remove(filePath)
	return
}

// GetFileByFileNamePathAndUserId 根据文件名、文件路径和用户ID查询文件
func (sf *sFile) GetFileByFileNamePathAndUserId(ctx context.Context, name, path string, userId uint) (file entity.File, err error) {
	err = dao.File.Ctx(ctx).Where(g.Map{
		dao.File.Columns().Name:   name,
		dao.File.Columns().Path:   path,
		dao.File.Columns().UserId: userId,
	}).Scan(&file)
	return
}

// Move 移动文件（fileId）到指定文件夹（dirId）下
func (sf *sFile) Move(ctx context.Context, userId uint, fileId, dirId string) error {
	dirPath, err := sf.GetFilePathByFileIdAndUserId(ctx, dirId, userId)
	if err != nil {
		return gerror.New("移动失败")
	}
	file, err := sf.GetFileByFileIdAndUserId(ctx, fileId, userId)
	src := sf.GetFilesRoot(ctx) + file.Path
	if err != nil {
		return gerror.New("移动失败")
	}
	_, err = dao.File.Ctx(ctx).Data(g.Map{
		dao.File.Columns().ParentId: dirId,
		dao.File.Columns().Path:     dirPath + "/" + file.Name,
	}).Where(g.Map{
		dao.File.Columns().Id:     fileId,
		dao.File.Columns().UserId: userId,
	}).Update()
	if err != nil {
		return gerror.New("移动失败")
	}
	absPath := sf.GetFilesRoot(ctx) + "/" + dirPath
	dest := absPath + "/" + file.Name
	// 若absPath不存在，则创建之
	isExists, _ := PathIsExists(absPath)
	if !isExists {
		os.Mkdir(absPath, os.ModePerm)
	}
	err = os.Rename(src, dest)
	if err != nil {
		return gerror.New("移动失败")
	}
	return nil
}

// PathIsExists 判断文件或文件夹是否在磁盘上
func PathIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	// 文件或文件夹存在
	if err == nil {
		return true, nil
	}
	// err 为文件或文件夹不存在的错误
	if os.IsNotExist(err) {
		return false, nil
	}
	// err 为其他错误，不确定文件或文件夹是否存在
	return false, err
}

func (sf *sFile) GetFileNameById(ctx context.Context, id string) (string, error) {
	val, err := dao.File.Ctx(ctx).Fields(dao.File.Columns().Name).Where(dao.File.Columns().Id, id).Value()
	if err != nil {
		return "", err
	}
	return val.String(), nil
}

func (sf *sFile) GetMultiFilesName(ctx context.Context, ids []string) (name string, err error) {
	if len(ids) == 1 {
		return sf.GetFileNameById(ctx, ids[0])
	} else {
		n, err := sf.GetFileNameById(ctx, ids[0])
		if err != nil {
			return n, err
		}
		m, err := sf.GetFileNameById(ctx, ids[1])
		if err != nil {
			return m, err
		}
		return n + "_" + m + "......", nil
	}
}