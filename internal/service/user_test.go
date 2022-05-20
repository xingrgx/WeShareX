package service_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/xingrgx/WeShareX/internal/service"
)

func Test_GetUser(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		fs, _ := service.User().GetUserByPassport(ctx, "xht")
		for _, f := range fs {
			t.Logf("f: %v\n", f)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		fs, _ := service.User().GetUserByNickname(ctx, "1")
		for _, f := range fs {
			t.Logf("f: %v\n", f)
		}
	})
}

func Test_GetNicknameById(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		nickname := service.User().GetNicknameById(ctx, 1111)
		t.Log(nickname)
	})
}