package service

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xingrgx/WeShareX/internal/model"
	"github.com/xingrgx/WeShareX/utility/response"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

// Middleware 获取中间件服务
func Middleware() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/login",
	}
}

// ResponseHandler 请求返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), err.Error())
		} else {

		}
	} else {
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), "", res)
		} else {
			// 什么都不做，业务API自行处理模板渲染的成功逻辑。
		}
	}
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	Context().Init(r, customCtx)
	if userEntity := Session().GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			Passport: userEntity.Passport,
			Nickname: userEntity.Nickname,
			Avatar: userEntity.Avatar,
			Email: userEntity.Email,
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
