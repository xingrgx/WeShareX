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
		Nickname: userEntity.Nickname,
		Avatar: userEntity.Avatar,
		Email: userEntity.Email,
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
		return err
	}
	if n > 0 {
		return gerror.Newf("您和账号 %s 无缘，换个新的吧~", passport)
	}
	return nil
}

// GetUserProfileByID 根据UID获取用户信息
func (su *sUser) GetUserProfileByID(ctx context.Context, userId uint) (user *model.UserProfileOutput, err error) {
	err = dao.User.Ctx(ctx).WherePri(userId).Scan(&user)
	// 查询失败
	if err != nil {
		return nil, err
	}
	// 查无此人
	if user == nil {
		return nil, nil
	}
	return
}

// UpdateProfileById 根据 UID 更新用户基本信息
func (su *sUser) UpdateProfileById(ctx context.Context, input model.UserProfileInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 获取上下文中的 UserId
		user := Context().Get(ctx).User
		userId := user.Id
		_, err := dao.User.Ctx(ctx).OmitEmpty().Data(input).Where(dao.User.Columns().Id, userId).Update()
		// 如果数据库更新成功，则更新session user 的信息
		if err == nil {
			sessionUser := Session().GetUser(ctx)
			sessionUser.Nickname = input.Nickname
			sessionUser.Gender = input.Gender
			err = Session().SetUser(ctx, sessionUser)
		}
		return err
	})
}

// UpdatePWD 用户修改密码服务
func (su *sUser) UpdatePWD(ctx context.Context, input model.UserPasswordInput) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		uid := Context().Get(ctx).User.Id
		n, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Password, input.OldPassword).
		Where(dao.User.Columns().Id, uid).Count()
		// 查询失败
		if err != nil {
			return err
		}
		// 查无此项
		if n == 0 {
			return gerror.New("原始密码错误！请重试")
		}
		_, err = dao.User.Ctx(ctx).Data(g.Map{
			dao.User.Columns().Password: input.NewPassword,
		}).Where(dao.User.Columns().Id, uid).Update()
		if err == nil {
			err = Session().RemoveUser(ctx)
		}
		return err
	})
}

// UpdateEmail 用户更改邮箱服务
func (su *sUser) UpdateEmail(ctx context.Context, newEmail string) error {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		uid := Context().Get(ctx).User.Id
		_, err := dao.User.Ctx(ctx).Data("email", newEmail).
		Where(dao.User.Columns().Id, uid).Update()
		if err == nil {
			sessionUser := Session().GetUser(ctx)
			sessionUser.Email = newEmail
			err = Session().SetUser(ctx, sessionUser)
		}
		return err
	})
}

func (su *sUser) GetUserByPassport(ctx context.Context, passport string) (users []*model.UserProfileOutput, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Passport, passport).Scan(&users)
	return
}

func (su *sUser) GetUserByNickname(ctx context.Context, nickname string) (users []*model.UserProfileOutput, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Nickname, nickname).Scan(&users)
	return
}

func (su *sUser) GetNicknameById(ctx context.Context, id uint) (nickname string) {
	val, _ := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Fields(dao.User.Columns().Nickname).Value()
	nickname = val.String()
	return
}