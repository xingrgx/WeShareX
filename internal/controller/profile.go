package controller

import (
	"context"

	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cProfile struct{}

var Profile cProfile

// IndexProfile 控制展示个人信息修改页面
func (cp *cProfile) IndexProfile(ctx context.Context, req *v1.IndexUpdateProfileReq) (res *v1.IndexUpdateProfileRes, err error) {
	uid := service.Context().Get(ctx).User.Id
	userPrf, err := service.User().GetUserProfileByID(ctx, uid)
	if err != nil {
		return nil, err
	} else {
		title := "用户基本资料"
		service.View().Render(ctx, model.View{
			Title: title,
			Data: userPrf,
		})
		return
	}
}

// UpdateProfile 控制更新用户基本信息
func (cp *cProfile) UpdateProfile(ctx context.Context, req *v1.UpdateProfileReq) (res *v1.UpdateProfileRes, err error) {
	err = service.User().UpdateProfileById(ctx, model.UserProfileInput {
		Id: req.Id,
		Nickname: req.Nickname,
		Gender: req.Gender,
	})
	return
}

// IndexChangePWD 控制展示更改密码页面
func (cp *cProfile) IndexUpdatePWD(ctx context.Context, req *v1.IndexUpdatePWDReq) (res *v1.IndexUpdatePWDRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "修改密码",
	})
	return
}

// UpdatePWD 控制修改密码
func (cp *cProfile) UpdatePWD(ctx context.Context, req *v1.UpdatePWDReq) (res *v1.UpdatePWDRes, err error) {
	err = service.User().UpdatePWD(ctx, model.UserPasswordInput {
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return
}

// IndexUpdateEmail 控制跳转至更改邮箱页面
func (cp *cProfile) IndexUpdateEmail(ctx context.Context, req *v1.IndexUpdateEmailReq) (res *v1.IndexUpdateEmailRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "更改邮箱",
	})
	return
}

// UpdateEmail 控制更改邮箱
func (cp *cProfile) UpdateEmail(ctx context.Context, req *v1.UpdateEmailReq) (res *v1.UpdateEmailRes, err error) {
	err = service.User().UpdateEmail(ctx, req.NewEmail)
	return
}