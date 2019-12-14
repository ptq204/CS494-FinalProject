package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	configs "final-project/configs"
	payload "final-project/server/action_payload"
	"final-project/server/constant"
	database "final-project/server/db/client"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net"
	"strconv"
)

func main() {
	config := &configs.SocketConfig{}
	err := configs.LoadConfigs()
	checkError(err)

	if err := viper.Unmarshal(config); err != nil {
		log.Fatal("load config: ", err)
	}
	var addr string
	addr = config.Host + ":" + strconv.Itoa(config.Port)
	fmt.Printf("Serve client on %s\n", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	database.GetConnectionDB()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 100)
	_, err := conn.Read(b[0:])
	resBuf := append(b[0:], 0)

	if err != nil {
		checkError(err)
	}
	for {
		if err == io.EOF {
			break
		}
		_, err = conn.Read(b[:])
		resBuf = append(resBuf, b[0:]...)
		if err != nil && err != io.EOF {
			checkError(err)
		}
	}

	action := binary.BigEndian.Uint32(resBuf[:4])
	tmpPayload := bytes.NewBuffer(resBuf[:4])

	d := gob.NewDecoder(tmpPayload)

	switch action {
	case constant.Login:
		fmt.Println("LOGINNN")
		var p payload.LoginPayload
		err := d.Decode(&p)
		if err != nil {
			checkError(err)
		}
		fmt.Printf("User %s login with password: %s\n", p.Username, p.Password)
		conn.Write([]byte("ACKKKK"))
	case constant.Chat:
		fmt.Println("CHATTTTT")
		var p payload.ChatPayload
		err := d.Decode(&p)
		if err != nil {
			checkError(err)
		}
		fmt.Printf("%s send msg to %s with content: %s\n", p.From, p.To, p.Message)
		conn.Write([]byte("ACKKKK"))
	default:
		fmt.Println(d)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
