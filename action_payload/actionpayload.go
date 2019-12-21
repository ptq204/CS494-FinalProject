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
	From      string   `json:"from"`
	To        []string `json:"to"`
	Message   string   `json:"message"`
	MultiUser int      `json:"multi_user"`
	Encrypt   int      `json:"encrypt"`
}

type UserPayload struct {
	Username string
}

type SetupUserPayload struct {
	Username string `json:"username"`
	NewInfo  string `json:"new_info"`
}

type DisconnectPayload struct {
}
