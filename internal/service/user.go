package service

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/model/entity"
	"github.com/xingrgx/WeShareX/internal/service/internal/dao"
)

type sUser struct{}

// User 获取用户服务
func User() *sUser {
	return &sUser{}
}

// Login 登录
func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) error {
	userEntity, err := s.GetUserByPassportAndPassword(ctx, in.Passport, in.Password)
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New("账号或密码错误！")
	}
	er := Session().SetUser(ctx, userEntity)
	if er != nil {
		return er
	}
	// 自动更新上线
	Context().SetUser(ctx, &model.ContextUser{
		Id:       userEntity.Id,
		Passport: userEntity.Passport,
	})
	return nil
}

// GetUserByPassportAndPassword 根据账号和密码（暂时明文）查询用户信息
func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Passport: passport,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}
