package constant

const (
	Login           = 1
	Chat            = 2
	Register        = 3
	Change_Password = 4
	Check_User      = 5
	Setup_Info      = 6
	Upload          = 7
	Download        = 8
	Chat_Multiple   = 9
)

type ActionDefine struct {
	code int
}

type ActionPayload interface {
}
