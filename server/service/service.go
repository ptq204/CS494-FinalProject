package service

import (
	payload "final-project/action_payload"
	"final-project/message"
	"final-project/server/business"
	"final-project/utils"
	"fmt"
	"net"
	"os"

	"golang.org/x/sync/syncmap"
)

func HandleLogin(c *net.Conn, resBuf []byte, clientConns *syncmap.Map) error {
	conn := *c
	fmt.Println("LOGINNN")
	var p payload.RegisterLoginPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	fmt.Printf("User %s login with password: %s\n", p.Username, p.Password)
	res := business.Signin(p.Username, p.Password, p.Encrypt)
	if res.ReturnCode == 1 {
		clientConns.Store(p.Username, *c)
	}
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleRegister(c *net.Conn, resBuf []byte) error {
	fmt.Println("REGISTERR")
	conn := *c
	var p payload.RegisterLoginPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		fmt.Printf("Error unmarshal: %s", err.Error())
		return err
	}
	fmt.Printf("User create new account with username: %s and password: %s\n", p.Username, p.Password)
	res := business.Register(p.Username, p.Password, p.Encrypt)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleChangePassword(c *net.Conn, resBuf []byte) error {
	fmt.Println("CHANGE_PASSWORD")
	conn := *c
	var p payload.ChangePasswordPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	fmt.Printf("User %s change password from %s to %s\n", p.Username, p.OldPassword, p.NewPassword)
	res := business.ChangePassword(p.Username, p.OldPassword, p.NewPassword, p.Encrypt)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleChat(c *net.Conn, resBuf []byte, clientConns *syncmap.Map) error {
	fmt.Println("CHATTTTT")
	conn := *c
	var p payload.ChatPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		fmt.Printf("Error unmarshal: %s", err.Error())
		return err
	}
	fmt.Printf("Message sent from: %s\n", p.To)
	fmt.Println(p.To)
	fmt.Printf("Content: %s\n", p.Message)
	for _, user := range p.To {
		toConn, ok := clientConns.Load(user)
		if ok {
			co := toConn.(net.Conn)
			res := message.ReturnMessageChat{From: p.From, To: user, Message: p.Message, Encrypt: p.Encrypt, ReturnCode: 1, ReturnMessage: ""}
			resBytes := utils.MarshalObject(res)
			co.Write(resBytes)
		}
	}
	conn.Write([]byte("ACKKKK"))
	return nil
}

func HandleFindUser(c *net.Conn, resBuf []byte) error {
	fmt.Println("FIND USER")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.FindUser(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleOnlineUser(c *net.Conn, resBuf []byte) error {
	fmt.Println("ONLINE USER")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.OnlineUser(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleUserBirthday(c *net.Conn, resBuf []byte) error {
	fmt.Println("USER BIRTHDAY")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.UserBirthday(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleUserName(c *net.Conn, resBuf []byte) error {
	fmt.Println("USER NAME")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.UserName(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleUserNote(c *net.Conn, resBuf []byte) error {
	fmt.Println("USER NOTE")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.UserNote(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleUserInfo(c *net.Conn, resBuf []byte) error {
	fmt.Println("USER INFO")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.UserInfo(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleSetupUserName(c *net.Conn, resBuf []byte) error {
	fmt.Println("SETUP NAME")
	conn := *c
	var p payload.SetupUserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.SetupName(p.Username, p.NewInfo)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleSetupUserDate(c *net.Conn, resBuf []byte) error {
	fmt.Println("SETUP DATE")
	conn := *c
	var p payload.SetupUserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.SetupDate(p.Username, p.NewInfo)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleSetupUserNote(c *net.Conn, resBuf []byte) error {
	fmt.Println("SETUP NOTE")
	conn := *c
	var p payload.SetupUserPayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := business.SetupNote(p.Username, p.NewInfo)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleUploadFile(c *net.Conn, resBuf []byte) error {
	conn := *c
	fmt.Println("UPLOAD FILE")
	var p payload.UploadFilePayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}
	res := message.ReturnMessageUpFile{ReturnCode: 1, ReturnMessage: "OK"}
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)

	fileName := p.FileName
	if p.AlterFileName != "" {
		fileName = p.AlterFileName
	}

	err = utils.ReceiveFile(c, fileName, p.FileSize, int(p.Encrypt))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("ENDDDDD UPLOAD")
	return err
}

func HandleDownloadFile(c *net.Conn, resBuf []byte) error {
	conn := *c
	fmt.Println("DOWNLOAD FILE")
	var p payload.DownloadFilePayload
	err := utils.UnmarshalObject(&p, resBuf)
	if err != nil {
		return err
	}

	checkDownFile := false
	var res message.ReturnMessageDownFile
	if _, err = os.Stat(p.FileName); os.IsNotExist(err) {
		res = message.ReturnMessageDownFile{
			ReturnCode:    2,
			ReturnMessage: "Requested file does not exist",
			FileName:      "",
			FileSize:      0,
		}
	} else {
		fi, _ := os.Stat(p.FileName)
		res = message.ReturnMessageDownFile{
			ReturnCode:    1,
			ReturnMessage: "OK",
			FileName:      fi.Name(),
			FileSize:      fi.Size(),
		}
		checkDownFile = true
	}
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)

	if checkDownFile {
		buff := make([]byte, 100)
		nBytes, err := conn.Read(buff)
		if string(buff[:nBytes]) == "OK" {
			err = utils.SendFileData(c, p.FileName, int(p.Encrypt))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	fmt.Println("ENDDDDD UPLOAD")
	return err
}
