package action_payload

type LoginPayload struct {
	Username string
	Password string
}

type ChatPayload struct {
	From    string
	To      string
	Message string
}
