package manager

import (
	"bytes"
	"encoding/binary"
	"final-project/configs"
	payload "final-project/server/action_payload"
	"final-project/utils"
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
	fmt.Println(tcpAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	clientService = ClientSocket{
		conn: conn,
	}

	return err
}

func (c *ClientSocket) GetConnection() net.Conn {
	return c.conn
}

func GetClientService() ClientSocket {
	return clientService
}

func (c *ClientSocket) SendDataRegisterLogin(actionType int, username string, password string) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.RegisterLoginPayload{Username: username, Password: password}
	buffPayload := utils.MarshalObject(&pl)

	fmt.Println("DATA LENGTHHHH: %s\n", len(buffPayload))

	buffDataLength := new(bytes.Buffer)
	dataLength := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLength, uint32(len(buffPayload)))
	err = binary.Write(buffDataLength, binary.BigEndian, dataLength)

	dataSend := make([]byte, len(buffPayload)+8)
	copy(dataSend[:4], buffDataLength.Bytes())
	copy(dataSend[4:8], buffAction.Bytes())
	copy(dataSend[8:], buffPayload)

	_, err = c.conn.Write([]byte(dataSend))

	return err
}

func (c *ClientSocket) SendDataChangePassword(actionType int, username string, oldPassword string, newPassword string) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.ChangePasswordPayload{OldPassword: oldPassword, NewPassword: newPassword}
	buffPayload := utils.MarshalObject(&pl)

	buffDataLength := new(bytes.Buffer)
	dataLength := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLength, uint32(len(buffPayload)))
	err = binary.Write(buffDataLength, binary.BigEndian, dataLength)

	dataSend := make([]byte, len(buffPayload)+8)
	copy(dataSend[:4], buffDataLength.Bytes())
	copy(dataSend[4:8], buffAction.Bytes())
	copy(dataSend[8:], buffPayload)

	_, err = c.conn.Write([]byte(dataSend))

	return err
}

func (c *ClientSocket) SendDataChat(actionType int, from string, to string, message string) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.ChatPayload{From: from, To: to, Message: message}
	buffPayload := utils.MarshalObject(&pl)

	buffDataLength := new(bytes.Buffer)
	dataLength := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLength, uint32(len(buffPayload)))
	err = binary.Write(buffDataLength, binary.BigEndian, dataLength)

	dataSend := make([]byte, len(buffPayload)+8)
	copy(dataSend[:4], buffDataLength.Bytes())
	copy(dataSend[4:8], buffAction.Bytes())
	copy(dataSend[8:], buffPayload)

	_, err = c.conn.Write([]byte(dataSend))

	return err
}

func (c *ClientSocket) ReadData() (string, error) {
	resByte, err := ioutil.ReadAll(c.conn)
	res := string(resByte[0:])
	return res, err
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
