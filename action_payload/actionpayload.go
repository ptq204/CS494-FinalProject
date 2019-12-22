package action_payload

type RegisterLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Encrypt  int32  `json:"encrypt"`
}

type ChangePasswordPayload struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Encrypt     int32  `json:"encrypt"`
}

type ChatPayload struct {
	From      string   `json:"from"`
	To        []string `json:"to"`
	Message   string   `json:"message"`
	MultiUser int32    `json:"multi_user"`
	Encrypt   int32    `json:"encrypt"`
}

type UserPayload struct {
	Username string
}

type UploadFilePayload struct {
	FileName      string `json:"file_name"`
	FileSize      int64  `json:"file_size"`
	Encrypt       int32  `json:"encrypt"`
	AlterFileName string `json:"alter_name"`
}

type DownloadFilePayload struct {
	FileName string `json:"file_name"`
	Encrypt  int32  `json:"encrypt"`
}

type SetupUserPayload struct {
	Username string `json:"username"`
	NewInfo  string `json:"new_info"`
}

type DisconnectPayload struct {
}
