package manager

import (
	"bytes"
	"encoding/binary"
	payload "final-project/action_payload"
	"final-project/utils"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type ClientSocket struct {
	conn         net.Conn
	currUsername string
}

type ClientManager struct{}

var clientService ClientSocket = ClientSocket{currUsername: ""}
var clientManager ClientManager

func GetManager() ClientManager {
	return clientManager
}

func Connect(ipServer string, portServer string) error {
	var addr string
	addr = ipServer + ":" + portServer
	fmt.Printf("Connect to server on %s\n", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	fmt.Println(tcpAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	clientService = ClientSocket{
		conn:         conn,
		currUsername: "",
	}

	return err
}

func (c *ClientSocket) GetConnection() net.Conn {
	return c.conn
}

func (c *ClientSocket) SetCurrUserName(username string) {
	c.currUsername = username
}

func (c *ClientSocket) GetCurrUserName() string {
	return c.currUsername
}

func GetClientService() ClientSocket {
	return clientService
}

func (c *ClientSocket) SendDataRegisterLogin(actionType int, username string, password string, encrypt int) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.RegisterLoginPayload{Username: username, Password: password, Encrypt: int32(encrypt)}
	buffPayload := utils.MarshalObject(&pl)

	fmt.Printf("DATA LENGTHHHH: %d\n", len(buffPayload))

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

func (c *ClientSocket) SendDataChangePassword(actionType int, username string, oldPassword string, newPassword string, encrypt int) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.ChangePasswordPayload{Username: username, OldPassword: oldPassword, NewPassword: newPassword, Encrypt: int32(encrypt)}
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

func (c *ClientSocket) SendDataChat(actionType int, from string, to []string, message string, multiUser int, encrypt int) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.ChatPayload{From: from, To: to, Message: message, MultiUser: int32(multiUser), Encrypt: int32(encrypt)}
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

func (c *ClientSocket) CheckUser(actionType int, username string) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.UserPayload{Username: username}
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

func (c *ClientSocket) SetupInfo(actionType int, username string, newInfo string) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.SetupUserPayload{Username: username, NewInfo: newInfo}
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

func (c *ClientSocket) SendFileMetada(actionType int, fileInfo os.FileInfo, alterFileName string, encrypt int) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.UploadFilePayload{FileName: fileInfo.Name(), FileSize: fileInfo.Size(), AlterFileName: alterFileName, Encrypt: int32(encrypt)}
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

func (c *ClientSocket) SendDownFileMetada(actionType int, fileName string, encrypt int) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.DownloadFilePayload{FileName: fileName, Encrypt: int32(encrypt)}
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

func (c *ClientSocket) Disconnect(actionType int) error {
	buffAction := new(bytes.Buffer)
	action := make([]byte, 4)
	binary.BigEndian.PutUint32(action, uint32(actionType))
	err := binary.Write(buffAction, binary.BigEndian, action)

	checkError(err)

	pl := payload.DisconnectPayload{}
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

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
