package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cAdmin struct {
}

var Admin cAdmin

func (ca *cAdmin) AdminIndex(ctx context.Context, req *v1.IndexAdminReq) (res *v1.IndexAdminRes, err error) {
	s := service.Session().GetUser(ctx).Status
	if s == 2 {
		service.View().Render(ctx, model.View{
			Title: "管理员",
			Error: "success",
		})
		return
	}
	service.View().RenderTpl(ctx, "index/admin/error.html", model.View{})
	return
}

func (ca *cAdmin) AdminIndexUsers(ctx context.Context, req *v1.AdminIndexUsersReq) (res *v1.AdminIndexUsersRes, err error) {
	s := service.Session().GetUser(ctx).Status
	users, _ := service.User().GetAllCommonUsers(ctx)
	if s == 2 {
		service.View().Render(ctx, model.View{
			Title: "管理员",
			Error: "success",
			Data: g.Map{
				"Users": users,
			},
		})
		return
	}
	service.View().RenderTpl(ctx, "index/admin/error.html", model.View{})
	return
}

func (ca *cAdmin) AdminEnable(ctx context.Context, req *v1.EnableReq) (res *v1.EnableRes, err error) {
	err = service.Admin().Enable(ctx, req.UserId)
	return
}

func (ca *cAdmin) AdminDisable(ctx context.Context, req *v1.DisableReq) (res *v1.DisableRes, err error) {
	err = service.Admin().Disable(ctx, req.UserId)
	return
}

func (ca *cAdmin) AdminProfile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	profile, err := service.Admin().GetCommonUserById(ctx, req.UserId)
	service.View().Render(ctx, model.View{
		Title: "修改信息",
		Data: g.Map{
			"User": profile,
		},
	})
	return
}

func (ca *cAdmin) AdminUpdate(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Admin().UpdateProfile(ctx, req.Input)
	return
}

