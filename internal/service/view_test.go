package service_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/xingrgx/WeShareX/internal/service"
)

func Test_GetBreadCrumbView(t *testing.T) {
	val := service.View().GetBreadCrumbView(ctx, "1vlyjjj5000cje8wloeexv8200doou52")
	g.Dump(val)
}
