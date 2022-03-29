package service

import (
	"context"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xingrgx/WeShareX/internal/model"
)

type sView struct{}

// View 获取视图服务
func View() *sView {
	return &sView{}
}

// RenderTpl 渲染指定模板页面
func (s *sView) RenderTpl(ctx context.Context, tpl string, data ...model.View) {
	var (
		request  = g.RequestFromCtx(ctx)
		viewObj  = model.View{}
		viewData = make(g.Map)
		stTitle  = g.Cfg().MustGet(ctx, `setting.title`).String()
	)
	// 获取前端传来的页面信息中的第一个
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = stTitle
	} else {
		viewObj.Title = viewObj.Title + " | " + stTitle
	}
	viewData = gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}
	// 首页中的登录的内置对象
	viewData["BuildIn"] = &viewBuildIn{httpRequest: request}
	// 内容模板
	if viewData["MainTpl"] == nil {
		viewData["MainTpl"] = s.getDefaultMainTpl(ctx)
	}
	_ = request.Response.WriteTpl(tpl, viewData)
}

// Render 渲染默认模板页面
func (s *sView) Render(ctx context.Context, data ...model.View) {
	s.RenderTpl(ctx, g.Cfg().MustGet(ctx, `viewer.indexLayout`).String(), data...)
}

// 获取视图存储目录
func (s *sView) getViewFolderName(ctx context.Context) string {
	return gstr.Split(g.Cfg().MustGet(ctx, "viewer.indexLayout").String(), "/")[0]
}

// 获取自动设置的MainTpl
func (s *sView) getDefaultMainTpl(ctx context.Context) string {
	var (
		viewFolderPrefix = s.getViewFolderName(ctx)
		urlPathArray     = gstr.SplitAndTrim(g.RequestFromCtx(ctx).URL.Path, "/")
		mainTpl          string
	)
	if len(urlPathArray) > 0 && urlPathArray[0] == viewFolderPrefix {
		urlPathArray = urlPathArray[1:]
	}
	switch {
	case len(urlPathArray) == 2:
		// 如果2级路由为数字，那么为模块的详情页面，那么路由固定为/xxx/detail。
		// 如果需要定制化内容模板，请在具体路由方法中设置MainTpl。
		if gstr.IsNumeric(urlPathArray[1]) {
			urlPathArray[1] = "detail"
		}
		mainTpl = viewFolderPrefix + "/" + gfile.Join(urlPathArray[0], urlPathArray[1]) + ".html"
	case len(urlPathArray) == 1:
		mainTpl = viewFolderPrefix + "/" + urlPathArray[0] + "/index.html"
	default:
		// 默认首页内容
		mainTpl = viewFolderPrefix + "/index/index.html"
	}
	return gstr.TrimLeft(mainTpl, "/")
}
