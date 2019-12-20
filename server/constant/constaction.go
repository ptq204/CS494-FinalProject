package constant

const (
	Login           = 1
	Chat            = 2
	Register        = 3
	Change_Password = 4
	UserInfo        = 5
	Setup_Info      = 6
	Upload          = 7
	Download        = 8
	Chat_Multiple   = 9
	UserNote        = 10
	UserBirthday    = 11
	UserOnline      = 12
	FindUser        = 13
	UserName        = 14
)

type ActionDefine struct {
	code int
}

type ActionPayload interface {
}
