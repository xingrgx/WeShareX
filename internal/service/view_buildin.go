package service

import "github.com/gogf/gf/v2/net/ghttp"

type viewBuildIn struct {
	httpRequest *ghttp.Request
}

func (s *viewBuildIn) UrlPath() string {
	return s.httpRequest.URL.Path
}