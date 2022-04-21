package service_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/xingrgx/WeShareX/internal/service"
)

var ctx context.Context

func Test_GetFileRelativePath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path, _ := service.File().GetFileRelativePath(ctx, 1, "root", "1.jpg")
		t.Assert(path, "/1.jpg")
	})
	gtest.C(t, func(t *gtest.T) {
		path, _ := service.File().GetFileRelativePath(ctx, 1, "1vlyjjj1440cj9tmkxq7omc200iktwak", "2.jpg")
		t.Assert(path, "/hello/world/2.jpg")
	})
}

func Test_CheckFileNameExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := service.File().CheckFileNameExist(ctx, "1.jpg", 1, "/hello/1.jpg")
		t.Assert(err, nil)
	})
}

// UNPASS
func Test_GetFilesRoot(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(service.File().GetFilesRoot(ctx), "files/1/root")
	})
}

func Test_GetFileByFileIdAndUserId(t *testing.T) {
	file, _ := service.File().GetFileByFileIdAndUserId(ctx, "1vlyjjj4uc0cje8jw4qq9q04b0i3o5gs", 1)
	g.Dump(file)
}

func Test_RenameFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := service.File().RenameFile(ctx, 1, "1vlyjjj4140cjan26rptr2s1005ro5sg", "张三")
		t.Assert(err, nil)
	})
}

func Test_GetFilePathByFileIdAndUserId(t *testing.T) {
	path, err := service.File().GetFilePathByFileIdAndUserId(ctx, "1vlyjjj6yo0cjbj71q7hhgs300drqvap", 1)
	if err != nil {
		g.Dump("获取失败")
	} else {
		g.Dump(path)
	}
}

func Test_GetFileByFileNamePathAndUserId(t *testing.T) {
	file, err := service.File().GetFileByFileNamePathAndUserId(ctx, "1.jpg", "/hi/1.jpg", 1)
	if err != nil {
		g.Dump("error")
	} else {
		g.Dump(file)
	}

}

func Test_GetDirFiles(t *testing.T) {
	m, _ := service.File().GetDirFiles(ctx, 1, "1vlyjjj4uc0cje8aq4455eo190skcs6g", 1, 10)
	g.Dump(m)
}
