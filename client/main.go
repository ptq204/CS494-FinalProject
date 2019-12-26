package main

import (
	"bufio"
	manager "final-project/client/manager"
	"final-project/client/service"
	"final-project/utils"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	checkExit := false
	for {
		fmt.Print("SocketApp~>: ")
		cmd, _ := reader.ReadString('\n')
		var commands []string
		commands, _ = utils.SplitCommand(cmd)

		if len(commands) == 4 && commands[0] == "connect" {
			ipServer := commands[1]
			portServer := commands[3]
			err := manager.Connect(ipServer, portServer)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				break
			}
		} else if strings.TrimRight(cmd, "\n") == "exit" {
			checkExit = true
			break
		} else {
			fmt.Println("Please connect before continuing")
		}
	}
	if !checkExit {
		clientService := manager.GetClientService()
		for {
			fmt.Print("SocketApp~>: ")
			cmd, _ := reader.ReadString('\n')
			cmd = strings.TrimRight(cmd, "\n")
			if !service.ParseCmdAndExecute(&clientService, cmd) {
				break
			}
		}
	} else {
		fmt.Println("Bye bye!!")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
