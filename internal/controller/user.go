package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cUser struct{}

var User = cUser{}

// Logout 用户注销控制
func (cu *cUser) Logout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	err = service.User().Logout(ctx)
	if err != nil {
		return
	}
	// 注销登录后重定向至登录页面
	loginUrl := service.Middleware().LoginUrl
	g.RequestFromCtx(ctx).Response.RedirectTo(loginUrl)
	return
}
