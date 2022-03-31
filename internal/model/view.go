package model

type View struct {
	Title    string
	Error    string
	MainTpl  string
	Redirect string
	Data     interface{}
}