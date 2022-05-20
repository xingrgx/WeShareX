package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
	return dao.Friends.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, e := dao.Friends.Ctx(ctx).Data(g.Map{
			dao.Friends.Columns().Me:     userId,
			dao.Friends.Columns().Friend: friendId,
			dao.Friends.Columns().Status: 1,
			dao.Friends.Columns().Time:   gtime.Now(),
		}).Save(); e != nil {
			return e
		}
		if _, e := dao.Friends.Ctx(ctx).Data(g.Map{
			dao.Friends.Columns().Me:     friendId,
			dao.Friends.Columns().Friend: userId,
			dao.Friends.Columns().Status: 0,
			dao.Friends.Columns().Time:   gtime.Now(),
		}).Save(); e != nil {
			return e
		}
		return nil
	})
}

func (sc *sChat) SetStatusTo2ById(ctx context.Context, userId, friendId uint) (err error) {
	return dao.Friends.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, err = dao.Friends.Ctx(ctx).Where(g.Map{
			dao.Friends.Columns().Me:     userId,
			dao.Friends.Columns().Friend: friendId,
		}).Data(dao.Friends.Columns().Status, 2).Update(); err != nil {
			return err
		}
		if _, err = dao.Friends.Ctx(ctx).Where(g.Map{
			dao.Friends.Columns().Me:     friendId,
			dao.Friends.Columns().Friend: userId,
		}).Data(dao.Friends.Columns().Status, 2).Update(); err != nil {
			return err
		}
		return nil
	})
}

func (sc *sChat) GetAllMsgs(ctx context.Context, userId, friendId uint) (msgs []*model.CMsg, err error) {
	var records []*entity.Record
	err = dao.Record.Ctx(ctx).Where(g.Map{
		dao.Record.Columns().SenderId:   userId,
		dao.Record.Columns().ReceiverId: friendId,
	}).WhereOr(g.Map{
		dao.Record.Columns().SenderId:   friendId,
		dao.Record.Columns().ReceiverId: userId,
	}).OrderAsc(dao.Record.Columns().Id).Scan(&records)
	gconv.Structs(records, &msgs)
	for _, msg := range msgs {
		msg.Sender = User().GetNicknameById(ctx, msg.SenderId)
		msg.Receiver = User().GetNicknameById(ctx, msg.ReceiverId)
	}
	return
}

func (sc *sChat) SaveMsg(ctx context.Context, senderId, receiverId uint, msg string) (err error) {
	_, err = dao.Record.Ctx(ctx).Data(g.Map {
		dao.Record.Columns().SenderId: senderId,
		dao.Record.Columns().ReceiverId: receiverId,
		dao.Record.Columns().Content: msg,
		dao.Record.Columns().Time: gtime.Now(),
	}).Save()
	return
}
