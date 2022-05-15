package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
		str  = gstr.TrimRight(input.FileIds, ",")
		ids  = gstr.Split(str, ",")
		name string
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
		s.Code = genCode()
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

// genCode 随机生成四位字符串
func genCode() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < 4; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func (ss *sShare) GetShareById(ctx context.Context, id string) (share *entity.Share, err error) {
	err = dao.Share.Ctx(ctx).Where(dao.File.Columns().Id, id).Scan(&share)
	return
}

func (ss *sShare) GetAllShares(ctx context.Context, userId uint) (shares []g.Map, err error) {
	var arr []entity.Share
	err = dao.Share.Ctx(ctx).Where(dao.Share.Columns().UserId, userId).Scan(&arr)
	for _, share := range arr {
		shares = append(shares, gconv.Map(share))
	}
	return
}

func (ss *sShare) GetShareByIdAndCode(ctx context.Context, id, code string) (share *entity.Share, err error) {
	err = dao.Share.Ctx(ctx).Where(g.Map {
		dao.Share.Columns().Id: id,
		dao.Share.Columns().Code: code,
	}).Scan(&share)
	return
}

// func (ss *sShare) Download(ctx context.Context, id, code string) (path string, err error) {
	
// }