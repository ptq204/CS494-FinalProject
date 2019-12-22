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
				if n == 3 && commands[1][1:] == "encrypt" {
					Login(commands[2], password, 1, clientService)
				} else if n == 2 {
					Login(commands[1], password, 0, clientService)
				}
			}
		case "register":
			if n != 2 && n != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'register' command")
			} else {
				password := utils.InputPassword()
				if n == 3 && commands[1] == "-encrypt" {
					Register(commands[2], password, 1, clientService)
				} else if n == 2 {
					Register(commands[1], password, 0, clientService)
				}
			}
		case "change_password":
			if n != 2 && n != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'change_password' command")
			} else {
				password := utils.InputPassword()
				newPassword := utils.InputNewPassword()
				if n == 3 && commands[1] == "-encrypt" {
					ChangePassword(commands[2], password, newPassword, 1, clientService)
				} else if n == 2 {
					ChangePassword(commands[1], password, newPassword, 0, clientService)
				}
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
		case "upload":
			files := []string{}
			if n == 2 {
				files = append(files, commands[1])
				UploadFile(files, "", "", clientService)
			} else if n == 3 && commands[1] == "-encrypt" {
				files = append(files, commands[2])
				UploadFile(files, "", "encrypt", clientService)
			} else if n == 4 && commands[1] == "-change_name" {
				files = append(files, commands[3])
				UploadFile(files, commands[2], "", clientService)
			} else if n >= 3 && commands[1] == "-multi_file" {
				for i := 2; i < n; i++ {
					files = append(files, commands[i])
				}
				UploadFile(files, "", "", clientService)
			} else {
				fmt.Println("(error) ERR wrong number of arguments for 'upload' command")
			}
		case "download":
			files := []string{}
			if n == 2 {
				files = append(files, commands[1])
				DownloadFile(files, "", clientService)
			} else if n == 3 && commands[1] == "-encrypt" {
				files = append(files, commands[2])
				DownloadFile(files, "encrypt", clientService)
			} else if n >= 3 && commands[1] == "-multi_file" {
				for i := 2; i < n; i++ {
					files = append(files, commands[i])
				}
				DownloadFile(files, "", clientService)
			} else {
				fmt.Println("(error) ERR wrong number of arguments for 'upload' command")
			}
		}
	}
	return true
}
