package controller

import (
	"context"

	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cProfile struct{}

var Profile cProfile

// Index 控制展示个人信息修改页面
func (cp *cProfile) Index(ctx context.Context, req *v1.IndexUpdateProfileReq) (res *v1.IndexUpdateProfileRes, err error) {
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

func (cp *cProfile) UpdateProfile(ctx context.Context, req *v1.UpdateProfileReq) (res *v1.UpdateProfileRes, err error) {
	err = service.User().UpdateProfileById(ctx, model.UserProfileInput {
		Id: req.Id,
		Nickname: req.Nickname,
		Gender: req.Gender,
	})
	return
}