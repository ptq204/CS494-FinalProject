package main

import (
	"final-project/client/cmd"
	"final-project/client/manager"
	"final-project/configs"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	config := &configs.SocketConfig{}
	err := configs.LoadConfigs()
	checkError(err)

	if err := viper.Unmarshal(config); err != nil {
		log.Fatal("load config: ", err)
	}
	err = manager.Connect(config)
	checkError(err)
	cmd.Execute()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		// panic(err)
	}
}
