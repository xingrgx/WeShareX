package controller

import (
	"context"

	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

// Login Controller

type cLogin struct{}

var Login = cLogin{}

// Index 控制首页跳转至登录页面
func (cl *cLogin) Index(ctx context.Context, req *v1.LoginIndexReq) (res *v1.LoginIndexRes, err error) {
	service.View().Render(ctx, model.View{})
	return
}

// Login 登录控制
func (cl *cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	err = service.User().Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})

	if err != nil {
		return
	} else {
		// 识别并跳转到登录前页面
		loginReferer := service.Session().GetLoginReferer(ctx)
		if loginReferer != "" {
			_ = service.Session().RemoveLoginReferer(ctx)
		}
		res.Referer = loginReferer
		return
	}
}
