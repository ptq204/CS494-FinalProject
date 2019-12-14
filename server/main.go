package main

import (
	configs "final-project/configs"
	"fmt"
	"net"
	"strconv"

	database "final-project/server/db/client"
	"final-project/server/db/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db.AutoMigrate(&entity.UserChannel{}, &entity.Message{}, &entity.File{}, &entity.Channel{}, &entity.User{})
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
	fmt.Println(conn)
	b := make([]byte, 100)
	_, err := conn.Read(b[0:])
	fmt.Println("CHeCKKKK")
	checkError(err)
	fmt.Println(string(b[0:]))
	conn.Write([]byte("ACKKKK"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
