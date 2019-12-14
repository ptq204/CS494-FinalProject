package message

const (
	Success = 1
	Unknown = -9999

	// Database error
	CannotCreateUser       = -101
	UsernameNotFound  = -102
	UserNotActive            = -103
	FilenameNotFound = -104
	CannotSetupUserInfo = -105
	CannotUploadFile     = -106
	CannotDownloadFile = -107


	// Unauthorized
	UserNotSignIn       = -200

	// BadRequest
	BusyUser       = -400
	UsernameExist = -401
	WrongPassword = -402
)

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

var (
	names = make(map[int32]string)
)

type ReturnMessage struct{
	ReturnCode int32
	ReturnMessage string
}

// GetErrorDecription return description for errorcode
func GetMessageDecription(code int32) string {
	if len(names) < 1 {
		names[Success] = "SUCCESS"
		names[Unknown] = "System Error. Please try again."

		names[CannotCreateUser] = "User cannot be created. Please try again."
		names[UsernameNotFound] = "Username not found. Please try again."
		names[UserNotActive] = "User is not active"
		names[FilenameNotFound] = "Filename not found. Please try again."
		names[CannotSetupUserInfo] = "Cannot set up user info. Please try again."

		names[UserNotSignIn] = "User not sigin."
		names[CannotUploadFile] = "Cannot upload file. Please try again."
		names[CannotDownloadFile] = "Cannot download file. Please ttry again."

		names[BusyUser] = "User is busy. Please try again."
		names[UsernameExist] = "Username is existed. Please try again."
		names[WrongPassword] = "Wrong password. Please try again."
	}

	return names[code]
}
