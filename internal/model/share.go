package model

type ShareInput struct {
	Id          string
	UserId      uint
	NickName    string
	FileIds     string
	NeverExpire bool
	ExpireAt  string
}

type ShareOutput struct {
	Url  string
	Code string
}
