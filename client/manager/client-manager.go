package manager

import (
	"final-project/configs"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
)

type ClientSocket struct {
	conn net.Conn
}

type ClientManager struct{}

var clientService ClientSocket
var clientManager ClientManager

func GetManager() ClientManager {
	return clientManager
}

func Connect(config *configs.SocketConfig) error {
	var addr string
	addr = config.Host + ":" + strconv.Itoa(config.Port)
	fmt.Printf("Connect to server on %s\n", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	clientService = ClientSocket{
		conn: conn,
	}

	return err
}

func GetClientService() ClientSocket {
	return clientService
}

func (c *ClientSocket) SendData(data string) error {
	_, err := c.conn.Write([]byte(data))
	return err
}

func (c *ClientSocket) ReadData() (string, error) {
	resByte, err := ioutil.ReadAll(c.conn)
	res := string(resByte[0:])
	return res, err
}
