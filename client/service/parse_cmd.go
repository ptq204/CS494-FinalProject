package service

import (
	"final-project/client/manager"
	"final-project/utils"
	"fmt"
)

func ParseCmdAndExecute(clientService *manager.ClientSocket, command string) bool {

	if clientService.GetCurrUserName() == "" {
		fmt.Println("Client havent connect yet")
	} else {
		fmt.Println(clientService.GetCurrUserName())
	}

	var commands []string
	commands, _ = utils.SplitCommand(command)

	if utils.IsExitCommand(commands[0]) {
		fmt.Println("Bye bye!!")
		Exit(clientService)
		return false
	}

	n := len(commands)

	if n > 0 {
		ctype := commands[0]
		switch ctype {
		case "login":
			if n != 2 && n != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'login' command")
			} else {
				password := utils.InputPassword()
				Login(commands[1], password, clientService)
			}
		case "register":
			if n != 2 && n != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'register' command")
			} else {
				password := utils.InputPassword()
				Register(commands[1], password, clientService)
			}
		case "change_password":
			if n != 2 && n != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'change_password' command")
			} else {
				password := utils.InputPassword()
				newPassword := utils.InputNewPassword()
				ChangePassword(commands[1], password, newPassword, clientService)
			}
		case "check_user":
			if n != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'check_user' command")
			} else {
				CheckUser(commands[2], commands[1][1:], clientService)
			}
		case "setup":
			if n != 4 {
				fmt.Println("(error) ERR wrong number of arguments for 'setup' command")
			} else {
				SetupUser(commands[3], commands[1][1:], commands[2], clientService)
			}
		case "chat":
			Chat(clientService)
		}
	}
	return true
}
