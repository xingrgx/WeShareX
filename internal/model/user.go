package model

// UserLoginInput 用户登录时输入的信息
type UserLoginInput struct {
	Passport string // 账号
	Password string // 密码（明文）
}

// UserRegisterInput 用户注册时输入的信息
type UserRegisterInput struct {
	Passport string // 账号
	Password string // 密码（明文）
	Nickname string // 昵称
	Email    string // 邮箱地址
}