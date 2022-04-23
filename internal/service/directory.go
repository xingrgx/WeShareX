package service

import (
	"archive/zip"
	"context"
	"io"
	"os"

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

func (sd *sDirectory) Zip(src []string, dest string) error {
	files := []*os.File{}
	for _, p := range src {
		file, _ := os.Open(p)
		files = append(files, file)
	}
	destZip, _ := os.Create(dest)
	defer destZip.Close()
	zipWriter := zip.NewWriter(destZip)
	defer zipWriter.Close()
	for _, file := range files {
		err := zipFile(file, "", zipWriter)
		if err != nil {
			return err
		}
	}
	return nil
}

func zipFile(file *os.File, prefix string, writer *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = zipFile(f, prefix, writer)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
