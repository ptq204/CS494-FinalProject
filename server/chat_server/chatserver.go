package chat_server

type ChatServer struct {
	Listen(address string) error
	BroadCast(command interface{}) error
	Start()
	Close()
}

	