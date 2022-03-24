package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/xingrgx/WeShareX/api/v1"
)

type cIndex struct{}

var Index = cIndex{}

func (i *cIndex) Index(ctx context.Context, req *v1.IndexReq) (res v1.IndexRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.GetView().Assign("tplMain", "index/include/main.html")
	_ = r.Response.WriteTpl("index/index.html")
	return
}