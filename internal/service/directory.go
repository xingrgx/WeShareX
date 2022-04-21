package service

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xingrgx/WeShareX/internal/model/entity"
)

type sDirectory struct{}

// 获取文件夹服务
func Directory() *sDirectory {
	return &sDirectory{}
}

// DeleteDirByDirIdAndUserId 删除文件夹
func (sd *sDirectory) DeleteDirByDirIdAndUserId(ctx context.Context, dirId string, userId uint) (err error) {
	err = deleteDirRecursively(ctx, dirId, userId)
	return
}

func deleteDirRecursively(ctx context.Context, dirId string, userId uint) (err error) {
	file, err := File().GetFileByFileIdAndUserId(ctx, dirId, userId)
	if err != nil {
		return
	}
	if file.Dir == 1 {
		filesMap, _ := File().GetDirFiles(ctx, userId, dirId, 0, 0)
		var files []*entity.File
		gconv.Structs(filesMap, &files)
		for _, file := range files {
			err = deleteDirRecursively(ctx, file.Id, userId)
			if err != nil {
				return
			}
		}
	}
	err = File().DeleteFileByFileIdAndUserId(ctx, file.Id, userId)
	if err != nil {
		return
	}
	return nil
}