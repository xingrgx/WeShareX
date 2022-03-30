package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
func (su *sUser) Login(ctx context.Context, in model.UserLoginInput) error {
	userEntity, err := su.GetUserByPassportAndPassword(ctx, in.Passport, in.Password)
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
func (su *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Passport: passport,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}

// Logout 用户注销登录
func (su *sUser) Logout(ctx context.Context) error {
	return Session().RemoveUser(ctx)
}

// Register 用户注册
func (su *sUser) Register(ctx context.Context, in model.UserRegisterInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var user *entity.User
		// 用户输入的注册信息和用户实体是否匹配
		if err := gconv.Struct(in, &user); err != nil {
			return err
		}
		// 检测用户名是否已存在
		if err := su.CheckPassportExist(ctx, user.Passport); err != nil {
			return err
		}
		// 新增用户数据
		_, err := dao.User.Ctx(ctx).Data(user).OmitEmpty().Save()
		return err
	})
}

// CheckPassportExist 检测用户是否已存在
func (su *sUser) CheckPassportExist(ctx context.Context, passport string) error {
	n, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Passport, passport).Count()
	if err != nil {
		return nil
	}
	if n > 0 {
		return gerror.Newf("您和账号 %s 无缘，换个新的吧~", passport)
	}
	return nil
}
