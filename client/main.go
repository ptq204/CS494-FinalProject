package main

import (
	configs "final-project/configs"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
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
	fmt.Printf("Connect to server on %s\n", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("QUYEN HELLO"))
	checkError(err)

	res, err := ioutil.ReadAll(conn)
	fmt.Println(string(res))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
