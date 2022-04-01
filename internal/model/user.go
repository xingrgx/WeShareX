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

// UserProfileOutput 用户个人展示的信息
type UserProfileOutput struct {
	Id       uint   // UID
	Nickname string // 昵称
	Avatar   string // 头像地址
	Gender   int    // 性别（0:未设置；1:男；2:女）
}

// UserProfileInput 用户更新的个人信息
type UserProfileInput struct {
	Id       uint   // UID
	Nickname string // 昵称
	Gender   int    // 性别（0:未设置；1:男；2:女）
}