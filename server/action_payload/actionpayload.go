package action_payload

type RegisterLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChatPayload struct {
	From    string
	To      string
	Message string
}
