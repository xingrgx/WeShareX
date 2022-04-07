package controller

import (
	"context"

	v1 "github.com/xingrgx/WeShareX/api/v1"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

type cIndex struct{}

var Index = cIndex{}

// Index 控制展示首页
func (i *cIndex) Index(ctx context.Context, req *v1.IndexReq) (res v1.IndexRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "首页",
	})
	return
}