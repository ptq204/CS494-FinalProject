package service

import (
	"bufio"
	"final-project/client/manager"
	"final-project/constant"
	"final-project/decrypt"
	e "final-project/encrypt"
	"final-project/message"
	"final-project/utils"
	"fmt"
	"net"
	"os"
	"strings"
)

func ChangePassword(username string, passStr string, newPassStr string, encrypt int, clientService *manager.ClientSocket) {
	fmt.Println("CHECK CHANGE PASSWORD")
	// clientService := manager.GetClientService()
	if encrypt == 1 {
		// Put encrypt function here
		username = e.Encrypt(constant.PASSPHRASE, username)
		passStr = e.Encrypt(constant.PASSPHRASE, passStr)
		newPassStr = e.Encrypt(constant.PASSPHRASE, newPassStr)
	}

	clientService.SendDataChangePassword(constant.Change_Password, username, passStr, newPassStr, encrypt)
	conn := clientService.GetConnection()
	var res message.ReturnMessage
	resData, _ := utils.ReadBytesResponse(&conn)
	err := utils.UnmarshalObject(&res, resData[:])
	if err != nil {
		fmt.Println("CANNOT UNMARSHAL")
		fmt.Println(err.Error())
		fmt.Println(string(resData[:]))
	}
	fmt.Printf("Return message: %d and %s\n", res.ReturnCode, res.ReturnMessage)
}

func Login(username string, password string, encrypt int, clientService *manager.ClientSocket) {
	fmt.Println("CHECK LOGINNNN")
	// clientService := manager.GetClientService()
	fmt.Println(&clientService)

	if encrypt == 1 {
		username = e.Encrypt(constant.PASSPHRASE, username)
		password = e.Encrypt(constant.PASSPHRASE, password)
	}
	clientService.SendDataRegisterLogin(constant.Login, username, password, encrypt)
	conn := clientService.GetConnection()
	fmt.Println(conn)
	// utils.TellReadDone(&conn)
	var res message.ReturnMessageLogin
	resData, _ := utils.ReadBytesResponse(&conn)
	err := utils.UnmarshalObject(&res, resData[:])
	if err != nil {
		fmt.Println("CANNOT UNMARSHAL")
		fmt.Println(err.Error())
		fmt.Println(string(resData[:]))
	}
	if res.ReturnCode == 1 {
		clientService.SetCurrUserName(username)
		// go listenMessageChat(conn)
	}
	fmt.Printf("Return message: %d and %s\n", res.ReturnCode, res.ReturnMessage)
}

func CheckUser(username string, flag string, clientService *manager.ClientSocket) {
	// clientService := manager.GetClientService()

	checkShowInfo := false

	if flag == "find" {
		clientService.CheckUser(constant.FindUser, username)
	} else if flag == "online" {
		clientService.CheckUser(constant.UserOnline, username)
	} else if flag == "show_date" {
		clientService.CheckUser(constant.UserBirthday, username)
	} else if flag == "show_note" {
		clientService.CheckUser(constant.UserNote, username)
	} else if flag == "show_name" {
		clientService.CheckUser(constant.UserName, username)
	} else if flag == "show" {
		clientService.CheckUser(constant.UserInfo, username)
		checkShowInfo = true
	}

	if !checkShowInfo {
		var res message.CheckUserResponse
		conn := clientService.GetConnection()
		resData, _ := utils.ReadBytesResponse(&conn)
		err := utils.UnmarshalObject(&res, resData[:])
		if err != nil {
			fmt.Println("CANNOT UNMARSHAL")
			fmt.Println(err.Error())
			fmt.Println(string(resData[:]))
		}
		fmt.Printf("Info: %s with %d and %s\n", res.Information, res.ReturnMessage.ReturnCode, res.ReturnMessage.ReturnMessage)
	} else {
		var res message.UserResponseInfo
		conn := clientService.GetConnection()
		resData, _ := utils.ReadBytesResponse(&conn)
		err := utils.UnmarshalObject(&res, resData[:])
		if err != nil {
			fmt.Println("CANNOT UNMARSHAL")
			fmt.Println(err.Error())
			fmt.Println(string(resData[:]))
		}
		fmt.Print("User info: ", res.User, " with ", res.ReturnCode, " and ", res.ReturnMessage, "\n")
	}
}

func Register(username string, password string, encrypt int, clientService *manager.ClientSocket) {
	// clientService := manager.GetClientService()

	if encrypt == 1 {
		username = e.Encrypt(constant.PASSPHRASE, username)
		password = e.Encrypt(constant.PASSPHRASE, password)
	}
	conn := clientService.GetConnection()
	clientService.SendDataRegisterLogin(constant.Register, username, password, encrypt)
	var res message.ReturnMessage
	resData, _ := utils.ReadBytesResponse(&conn)
	err := utils.UnmarshalObject(&res, resData[:])
	if err != nil {
		fmt.Println("CANNOT UNMARSHAL")
		fmt.Println(err.Error())
		fmt.Println(string(resData[:]))
	}
	fmt.Printf("Return message: %d and %s\n", res.ReturnCode, res.ReturnMessage)
}

func SetupUser(username string, flag string, newInfo string, clientService *manager.ClientSocket) {
	// clientService := manager.GetClientService()
	fmt.Println(&clientService)

	if flag == "name" {
		clientService.SetupInfo(constant.SetupName, username, newInfo)
	} else if flag == "date" {
		clientService.SetupInfo(constant.SetupDate, username, newInfo)
	} else if flag == "note" {
		clientService.SetupInfo(constant.SetupNote, username, newInfo)
	}

	conn := clientService.GetConnection()
	var res message.ReturnMessage
	resData, _ := utils.ReadBytesResponse(&conn)
	err := utils.UnmarshalObject(&res, resData[:])
	if err != nil {
		fmt.Println("CANNOT UNMARSHAL")
		fmt.Println(err.Error())
		fmt.Println(string(resData[:]))
	}
	fmt.Printf("Return message: %d and %s\n", res.ReturnCode, res.ReturnMessage)
}

func Chat(clientService *manager.ClientSocket) {
	// clientService := manager.GetClientService()
	currUserName := clientService.GetCurrUserName()

	conn := clientService.GetConnection()
	go listenMessageChat(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		users := []string{}
		fmt.Print("Chat~>: ")
		cmd, _ := reader.ReadString('\n')

		commands, _ := utils.SplitCommand(cmd)
		n := len(commands)

		multiUser := 0
		encrypt := 0

		check := false

		if n >= 2 {
			if n >= 3 {
				if (commands[1][1:] == "encrypt" && commands[2][1:] == "multi_user") || (commands[1][1:] == "multi_user" && commands[2][1:] == "encode") {
					for i := 3; i < n; i++ {
						users = append(users, commands[i])
					}
					multiUser = 1
					encrypt = 1
					check = true
				} else if commands[1] == "encrypt" {
					for i := 2; i < n; i++ {
						users = append(users, commands[i])
					}
					multiUser = 0
					encrypt = 1
					check = true
				} else if commands[1] == "multi_user" {
					for i := 2; i < n; i++ {
						users = append(users, commands[i])
					}
					multiUser = 1
					encrypt = 0
					check = true
				} else {
					for i := 1; i < n; i++ {
						users = append(users, commands[i])
					}
					multiUser = 0
					encrypt = 0
					check = true
				}
			} else if n == 2 {
				users = append(users, commands[1])
				check = true
			}

			if commands[0] == "chat" && check {
				fmt.Print(">> Me: ")
				msg, _ := reader.ReadString('\n')
				msg = strings.TrimRight(msg, "\n")
				if encrypt == 1 {
					// Put encrypt function here
					msg = e.Encrypt(constant.PASSPHRASE, msg)
				}
				clientService.SendDataChat(constant.Chat, currUserName, users, msg, multiUser, encrypt)
			}
		}
	}
}

func UploadFile(fileNames []string, alterFileName string, flag string, clientService *manager.ClientSocket) {

	encrypt := 0
	if flag == "encrypt" {
		encrypt = 1
	}

	for _, fileName := range fileNames {
		fi, err := os.Stat(fileName)
		conn := clientService.GetConnection()
		if err != nil {
			fmt.Println("Cannot process choosen file")
			return
		}
		clientService.SendFileMetada(constant.Upload, fi, alterFileName, encrypt)

		var resUpFile message.ReturnMessageUpFile
		resData, _ := utils.ReadBytesResponse(&conn)
		err = utils.UnmarshalObject(&resUpFile, resData[:])
		if err != nil {
			fmt.Println("CANNOT UNMARSHAL")
			fmt.Println(err.Error())
			fmt.Println(string(resData[:]))
		} else if resUpFile.ReturnCode == 1 && resUpFile.ReturnMessage == "OK" {
			err = utils.SendFileData(&conn, fi.Name(), encrypt)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func DownloadFile(fileNames []string, flag string, clientService *manager.ClientSocket) {
	encrypt := 0
	if flag == "encrypt" {
		encrypt = 1
	}

	for _, fileName := range fileNames {
		conn := clientService.GetConnection()
		clientService.SendDownFileMetada(constant.Download, fileName, encrypt)

		var resDownFile message.ReturnMessageDownFile
		resData, _ := utils.ReadBytesResponse(&conn)
		err := utils.UnmarshalObject(&resDownFile, resData[:])
		if err != nil {
			fmt.Println("CANNOT UNMARSHAL")
			fmt.Println(err.Error())
			fmt.Println(string(resData[:]))
		} else if resDownFile.ReturnCode == 1 && resDownFile.ReturnMessage == "OK" {
			conn.Write([]byte("OK"))
			err = utils.ReceiveFile(&conn, resDownFile.FileName, resDownFile.FileSize, encrypt)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			conn.Write([]byte("CONTINUTE"))
		}
	}
}

func Exit(clientService *manager.ClientSocket) {
	// clientService := manager.GetClientService()
	clientService.Disconnect(constant.Exit)
}

func listenMessageChat(conn net.Conn) {
	for {
		var res message.ReturnMessageChat
		resData, _ := utils.ReadBytesResponse(&conn)
		err := utils.UnmarshalObject(&res, resData[:])
		if err != nil {
			continue
		}
		// Check if message is encrypted
		// if yes, decrypt it
		if res.Encrypt == 1 {
			message := decrypt.Decrypt(constant.PASSPHRASE, res.Message)
			fmt.Printf("%s : %s\n", res.From, string(message))
		} else {
			fmt.Printf("%s : %s\n", res.From, string(res.Message))
		}

	}
}
