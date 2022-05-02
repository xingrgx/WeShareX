package controller

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
	"github.com/xingrgx/WeShareX/utility/response"
)

type cShare struct{}

var Share cShare

func (cs *cShare) ShareToCreate(ctx context.Context, req *v1.ShareToCreateReq) (res *v1.ShareToCreateRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "分享页面",
	})
	return
}

func (cs *cShare) ShareCreate(ctx context.Context, req *v1.ShareCreateReq) (res *v1.ShareCreateRes, err error) {
	b, _ := strconv.ParseBool(req.NeverExpire)
	id := guid.S()
	input := model.ShareInput{
		Id:          id,
		UserId:      service.Session().GetUser(ctx).Id,
		NickName:    service.Session().GetUser(ctx).Nickname,
		FileIds:     req.FileIds,
		NeverExpire: b,
		ExpireAt:    req.ExpireTime,
	}
	err = service.Share().CreateShare(ctx, input)
	if err != nil {
		return
	}
	share, err := service.Share().GetShareById(ctx, id)
	var output model.ShareOutput
	output.Url = "share/get?id=" + id
	output.Code = share.Code
	response.Json(g.RequestFromCtx(ctx), 0, "", output)
	return
}

func (cs *cShare) ShareList(ctx context.Context, req *v1.ShareListReq) (res *v1.ShareListRes, err error) {
	sharesMap, err := service.Share().GetAllShares(ctx, service.Session().GetUser(ctx).Id)
	service.View().Render(ctx, model.View{
		Title: "我的分享",
		Data: g.Map{
			"sharesMap": sharesMap,
		},
	})
	return
}

func (cs *cShare) ShareLink(ctx context.Context, req *v1.ShareLinkReq) (res *v1.ShareCreateRes, err error) {
	share, err := service.Share().GetShareById(ctx, req.Id)
	if err != nil {
		response.Json(g.RequestFromCtx(ctx), 500, "获取错误")
	}
	var output model.ShareOutput
	output.Url = "share/get?id=" + share.Id
	output.Code = share.Code
	response.Json(g.RequestFromCtx(ctx), 0, "", output)
	return
}