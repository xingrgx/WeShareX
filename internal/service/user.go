package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/model/entity"
	"github.com/xingrgx/WeShareX/internal/service/internal/dao"
)

type sUser struct{}

// User 用户服务
func User() *sUser {
	return &sUser{}
}

// Login 执行登录
func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (*entity.User, error) {
	return s.GetUserByPassportAndPassword(ctx, in.Passport, in.Password)
}

// GetUserByPassportAndPassword 根据账号和密码（暂时明文）查询用户信息
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Passport: passport,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}
