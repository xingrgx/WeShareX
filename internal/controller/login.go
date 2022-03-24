package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

// Login Controller

type cLogin struct{}

var Login = cLogin{}

func (cl *cLogin) Index(ctx context.Context, req *v1.LoginIndexReq) (res *v1.LoginIndexRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteTpl("index/login/index.html")
	return
}

func (cl *cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	user, err := service.User().Login(ctx, model.UserLoginInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	r := g.RequestFromCtx(ctx)
	if user != nil {
		r.Response.Writeln("Login Successed")
	} else {
		r.Response.Writeln("Login Failed")
	}
	return
}
