package action_payload

type RegisterLoginPayload struct {
	Username string
	Password string
}

type ChatPayload struct {
	From    string
	To      string
	Message string
}
