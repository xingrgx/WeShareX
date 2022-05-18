package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/model/entity"
	"github.com/xingrgx/WeShareX/internal/service/internal/dao"
)

type sChat struct{}

func Chat() *sChat {
	return &sChat{}
}

func (sc *sChat) GetAllFriends(id uint) (friends []*model.FriendProfile, err error) {
	err = g.Model("wsx_user u").LeftJoin("wsx_friends f", "u.id=f.friend").
		Fields("u.id, u.passport, u.nickname, f.status").Where("f.me", id).Scan(&friends)
	return
}

func (sc *sChat) AddFriend(ctx context.Context, userId, friendId uint) (err error) {
	var me, friend *entity.Friends
	if err = gconv.Struct(g.Map{
		"Me":     userId,
		"Friend": friendId,
	}, &me); err != nil {
		return err
	}
	if err = gconv.Struct(g.Map{
		"Me":     friendId,
		"Friend": userId,
	}, &friend); err != nil {
		return err
	}
	return dao.Friends.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, e := dao.Friends.Ctx(ctx).Data(me).Save(); e != nil {
			return e
		}
		if _, e := dao.Friends.Ctx(ctx).Data(friend).Save(); e != nil {
			return e
		}
		return nil
	})
}
