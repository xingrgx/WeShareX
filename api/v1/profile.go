package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// IndexUpdateProfileReq 跳转至个人信息修改页面的请求接口
type IndexUpdateProfileReq struct {
	g.Meta `path:"/profile" method:"get" summary:"展示个人信息修改页面" tags:"修改个人信息"`
}

// IndexUpdateProfileRes 跳转至个人信息修改页面的响应接口
type IndexUpdateProfileRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// UpdateProfileReq 个人信息修改请求接口
type UpdateProfileReq struct {
	g.Meta   `path:"/profile" method:"post" summary:"修改个人基本信息" tags:"修改个人信息"`
	Id       uint   `json:"id"       dc:"用户ID"`
	Nickname string `json:"nickname" dc:"昵称"`
	Gender   int    `json:"gender"   dc:"性别（0:未设置；1:男；2:女）"`
}

// UpdateProfileRes 个人信息修改响应接口
type UpdateProfileRes struct {
}

// IndexUpdatePasswordReq 跳转至用户修改密码页面的请求接口
type IndexUpdatePWDReq struct {
	g.Meta `path:"/profile/password" method:"get" summary:"展示修改密码页面" tags:"修改密码"`
}

// IndexUpdatePasswordRes 跳转至用户修改密码页面的响应接口
type IndexUpdatePWDRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// UpdatePWDReq 修改密码的请求接口
type UpdatePWDReq struct {
	g.Meta `path:"/profile/password" method:"post" summary:"修改密码" tags:"个人"`
	OldPassword string `json:"oldPassword" dc:"旧密码"`
	NewPassword string `json:"newPassword" dc:"新密码"`
}

// UpdatePWDRes 修改密码的响应接口
type UpdatePWDRes struct {

}

// IndexUpdateEmailReq 跳转至更改邮箱页面的请求接口
type IndexUpdateEmailReq struct {
	g.Meta `path:"/profile/email" method:"get" summary:"展示更改邮箱页面" tags:"个人"`
}

// IndexUpdateEmailRes 跳转至更改邮箱页面的响应接口
type IndexUpdateEmailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

// UpdateEmailReq 更改邮箱的请求接口
type UpdateEmailReq struct {
	g.Meta `path:"/profile/email" method:"post" summary:"更改邮箱" tags:"个人"`
	NewEmail string `json:"newEmail" dc:"新邮箱"`
}

// UpdateEmailRes 更改邮箱的响应接口
type UpdateEmailRes struct {

}