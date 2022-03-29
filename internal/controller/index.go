package controller

import (
	"context"

	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cIndex struct{}

var Index = cIndex{}

func (i *cIndex) Index(ctx context.Context, req *v1.IndexReq) (res v1.IndexRes, err error) {
	//r := g.RequestFromCtx(ctx)
	//r.Response.WriteTpl("index/index.html")
	service.View().Render(ctx, model.View{
		Title: "首页",
	})
	return
}