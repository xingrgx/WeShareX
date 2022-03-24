package model

// UserLoginInput 用户登录时输入的信息
type UserLoginInput struct {
	Passport string // 账号
	Password string // 密码（明文）
}