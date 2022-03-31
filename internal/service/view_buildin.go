package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xingrgx/WeShareX/internal/consts"
)

type viewBuildIn struct {
	httpRequest *ghttp.Request
}

func (s *viewBuildIn) UrlPath() string {
	return s.httpRequest.URL.Path
}

// 根据性别字段内容返回性别
func (s *viewBuildIn) Gender(gender int) string {
	switch gender {
	case consts.UserGenderMale:
		return "男"
	case consts.UserGenderFemale:
		return "女"
	default:
		return "未知"
	}
}
