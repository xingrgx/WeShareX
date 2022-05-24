package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/xingrgx/WeShareX/internal/controller"
	"github.com/xingrgx/WeShareX/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Index,
					controller.Login, // 登录
					controller.User,
					controller.Register,
					controller.Profile,
					controller.File,
					controller.Directory,
					controller.Share,
					controller.Chat,
					controller.Admin,
				)
			})
			s.Run()
			return nil
		},
	}
)
