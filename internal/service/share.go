package service

import (
	"context"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/model/entity"
	"github.com/xingrgx/WeShareX/internal/service/internal/dao"
)

type sShare struct{}

func Share() *sShare {
	return &sShare{}
}

func (ss *sShare) CreateShare(ctx context.Context, input model.ShareInput) error {
	var (
		str   = gstr.TrimRight(input.FileIds, ",")
		ids   = gstr.Split(str, ",")
		name  string
	)

	name, err := File().GetMultiFilesName(ctx, ids)
	if err != nil {
		return err
	}

	return dao.Share.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var s *entity.Share
		if err := gconv.Struct(input, &s); err != nil {
			return err
		}
		s.Name = name
		if _, err := dao.Share.Ctx(ctx).Data(s).Save(); err != nil {
			return err
		}
		for _, id := range ids {
			if _, err := dao.Relation.Ctx(ctx).Data(g.Map{
				dao.Relation.Columns().Id:      guid.S(),
				dao.Relation.Columns().ShareId: input.Id,
				dao.Relation.Columns().FileId:  id,
			}).Save(); err != nil {
				return err
			}
		}
		return nil
	})
}
