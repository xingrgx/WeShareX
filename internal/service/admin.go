package service

import (
	"context"

	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/model/entity"
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

func (sa *sAdmin) GetCommonUserById(ctx context.Context, userId uint) (output *model.UpdateProfile, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userId).Scan(&output)
	return
}

func (sa *sAdmin) UpdateProfile(ctx context.Context, input entity.User) (err error) {
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, input.Id).Data(input).OmitNil().Update()
	return
}
