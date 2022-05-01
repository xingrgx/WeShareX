package model

type ShareInput struct {
	Id          string
	UserId      uint
	FileIds     string
	NeverExpire bool
	ExpireTime  string
}

type ShareOutput struct {
	Url  string
	Code string
}