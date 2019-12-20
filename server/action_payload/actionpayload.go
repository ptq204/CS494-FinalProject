package action_payload

type RegisterLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChatPayload struct {
	From    string
	To      string
	Message string
}
