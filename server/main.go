package main

import (
	configs "final-project/configs"
	payload "final-project/server/action_payload"
	"final-project/server/business"
	"final-project/server/constant"
	database "final-project/server/db/client"
	entity "final-project/server/db/entity"
	"final-project/utils"
	"fmt"
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	db := database.GetConnectionDB()
	db.AutoMigrate(&entity.User{}, &entity.Channel{}, &entity.Message{}, &entity.File{}, &entity.UserChannel{})
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

	resBuf, action, err := utils.ReadBytesData(&conn)
	checkError(err)

	fmt.Println("Action type: %d\n", action)

	switch action {
	case constant.Login:
		fmt.Println("LOGINNN")
		var p payload.RegisterLoginPayload
		err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
		if err != nil {
			checkError(err)
		}
		fmt.Printf("User %s login with password: %s\n", p.Username, p.Password)
		res := business.Signin(p.Username, p.Password)
		resBytes := utils.MarshalObject(res)
		conn.Write(resBytes)
		break
	case constant.Register:
		fmt.Println("REGISTERR")
		var p payload.RegisterLoginPayload
		err := utils.UnmarshalObject(&p, resBuf[4:])
		if err != nil {
			checkError(err)
		}
		fmt.Printf("User %s register with password: %s\n", p.Username, p.Password)
		conn.Write([]byte("ACKKKK"))
	case constant.Chat:
		fmt.Println("CHATTTTT")
		var p payload.ChatPayload
		err := utils.UnmarshalObject(&p, resBuf[4:])
		if err != nil {
			checkError(err)
		}
		fmt.Printf("%s send msg to %s with content: %s\n", p.From, p.To, p.Message)
		conn.Write([]byte("ACKKKK"))
	default:
		fmt.Println("Default")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
