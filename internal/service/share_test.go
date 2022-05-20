package service

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/xingrgx/WeShareX/internal/model"
)

func Test_CreateShare(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := Share().CreateShare(ctx, model.ShareInput{
			Id:          guid.S(),
			UserId:      1,
			FileIds:     "1vlyjjjevw0cjjdwkhwmv80100rp1d0b,1vlyjjjevw0cjjdwkipfus4500uw3o6r,1vlyjjjevw0cjjdwno7m748w00sth5ls,",
			NeverExpire: false,
			ExpireAt:  "2022-06-01 09:37:44",
		})
		gtest.AssertNil(err)
	})
}

func Test_GetShareById(t *testing.T) {
	s, _ := Share().GetShareById(ctx, "18yz38e62o0cjo2hx6t11c4100xxj35n")
	t.Log(s)
}

func Test_genCode(t *testing.T) {
	t.Log(genCode())
	t.Log(genCode())
	t.Log(genCode())
	t.Log(genCode())
	t.Log(genCode())
}

func Test_GetAllShares(t *testing.T) {
	shares, _ := Share().GetAllShares(ctx, 1)
	t.Log(shares)
}

func Test_Download(t *testing.T) {
	Share().Download(ctx, "18yz38e4xw0cjoc9p2xbxsw100g3yl7d", "eMWk")
}