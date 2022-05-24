package service

import (
	"context"

	"github.com/xingrgx/WeShareX/internal/service/internal/dao"
)

type sAdmin struct{}

func Admin() *sAdmin {
	return &sAdmin{}
}

func (sa *sAdmin) Enable(ctx context.Context, userId uint) (err error) {
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userId).Data(dao.User.Columns().Status, 1).Update()
	return
}

func (sa *sAdmin) Disable(ctx context.Context, userId uint) (err error) {
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userId).Data(dao.User.Columns().Status, 0).Update()
	return
}