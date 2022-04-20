package model

// 文件上传时的文件信息
type FileUploadInput struct {
	Id       string
	UserId   uint
	Name     string
	ParentId string
	Dir      int
	Type     string
	Path     string
	Size     int64
}

type FileCrumb struct {
	Id string
	Name string
	ParentId string
}
