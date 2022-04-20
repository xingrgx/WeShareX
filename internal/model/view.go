package model

type View struct {
	Title    string
	Error    string
	MainTpl  string
	Redirect string
	Data     interface{}
	BreadCrumbs []BreadCrumbView // 路径的面包屑视图
}

// BreadCrumbView 面包屑视图结构
type BreadCrumbView struct {
	Name string
	Url  string
	CurrentPathId string // 即当前文件夹的id 
}
