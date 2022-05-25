package service_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/xingrgx/WeShareX/internal/service"
)

func Test_GetCommonUserById(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		output, _ := service.Admin().GetCommonUserById(ctx, 2)
		t.Log(output)
	})
}
