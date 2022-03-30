package controller

import (
	"context"

	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

// Register Controller
type cRegister struct{}

var Register = cRegister{}

// Index 控制首页跳转至注册页面
func (cr *cRegister) Index(ctx context.Context, req *v1.RegisterIndexReq) (res *v1.RegisterIndexRes, err error) {
	service.View().Render(ctx, model.View{})
	return
}

// Register 控制注册
func (cr *cRegister) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	res = &v1.RegisterRes{}
	err = service.User().Register(ctx, model.UserRegisterInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
	})
	return
}
