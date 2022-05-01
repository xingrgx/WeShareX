package service_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/internal/service"
)

func Test_CreateShare(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := service.Share().CreateShare(ctx, model.ShareInput{
			Id:          guid.S(),
			UserId:      1,
			FileIds:     "1vlyjjjevw0cjjdwkhwmv80100rp1d0b,1vlyjjjevw0cjjdwkipfus4500uw3o6r,1vlyjjjevw0cjjdwno7m748w00sth5ls,",
			NeverExpire: false,
			ExpireTime:  "2022-06-01 09:37:44",
		})
		gtest.AssertNil(err)
	})
}
