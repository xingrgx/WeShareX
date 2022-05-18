package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xingrgx/WeShareX/internal/model"
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