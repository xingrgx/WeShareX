package service_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/xingrgx/WeShareX/internal/service"
)

func Test_GetAllFriends(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		fs, _ := service.Chat().GetAllFriends(1)
		for _, f := range fs {
			t.Logf("f: %v\n", f)
		}
	})
}

func Test_AddFrined(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		e := service.Chat().AddFriend(ctx, 1, 25)
		t.AssertNil(e)
	})
}

func Test_GetAllMsgs(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, _ := service.Chat().GetAllMsgs(ctx, 1, 22)
		for i, r := range s {
			t.Log(i, r)
		}
		
	})
}