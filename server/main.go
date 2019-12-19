package main

import (
	configs "final-project/configs"
	"final-project/server/constant"
	database "final-project/server/db/client"
	entity "final-project/server/db/entity"
	service "final-project/server/service"
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
		service.HandleLogin(&conn, resBuf)
		break
	case constant.Register:
		service.HandleRegister(&conn, resBuf)
		break
	case constant.Change_Password:
		service.HandleChangePassword(&conn, resBuf)
		break
	case constant.Chat:
		service.HandleChat(&conn, resBuf)
		break
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
