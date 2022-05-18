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
