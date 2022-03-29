package service

import (
	"context"

	"github.com/xingrgx/WeShareX/internal/model/entity"
)

type sSession struct{}

const (
	sessionKeyUser         = "SessionKeyUser"    // 用户信息存放在Session中的Key
	sessionKeyLoginReferer = "SessionKeyReferer" // Referer存储，当已存在该session时不会更新。用于用户未登录时引导用户登录，并在登录后跳转到登录前页面。
	sessionKeyNotice       = "SessionKeyNotice"  // 存放在Session中的提示信息，往往使用后则删除
)

// Session 获取Session服务
func Session() *sSession {
	return &sSession{}
}

// SetUser 设置用户Session
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return Context().Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.User{}
}

// RemoveUser 删除用户Session
func (s *sSession) RemoveUser(ctx context.Context) error {
	myCtx := Context().Get(ctx)
	if myCtx != nil {
		return myCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}

// 设置LoginReferer.
func (s *sSession) SetLoginReferer(ctx context.Context, referer string) error {
	if s.GetLoginReferer(ctx) == "" {
		customCtx := Context().Get(ctx)
		if customCtx != nil {
			return customCtx.Session.Set(sessionKeyLoginReferer, referer)
		}
	}
	return nil
}

// 获取LoginReferer.
func (s *sSession) GetLoginReferer(ctx context.Context) string {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.MustGet(sessionKeyLoginReferer).String()
	}
	return ""
}

// 删除LoginReferer.
func (s *sSession) RemoveLoginReferer(ctx context.Context) error {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyLoginReferer)
	}
	return nil
}