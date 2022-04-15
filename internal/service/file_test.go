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

func Test_GetRootFiles(t *testing.T) {
	files, _ := service.File().GetRootFiles(ctx, 1, 1, 10)
	g.Dump(files)
	g.Dump(len(files))
}

func Test_CountRootFiles(t *testing.T) {
	count, _ := service.File().CountRootFiles(ctx, 1)
	g.Dump(count)
}