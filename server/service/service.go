package service

import (
	payload "final-project/server/action_payload"
	"final-project/server/business"
	"final-project/utils"
	"fmt"
	"net"
)

func HandleLogin(c *net.Conn, resBuf []byte) error {
	conn := *c
	fmt.Println("LOGINNN")
	var p payload.RegisterLoginPayload
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
	if err != nil {
		return err
	}
	fmt.Printf("User %s login with password: %s\n", p.Username, p.Password)
	res := business.Signin(p.Username, p.Password)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleRegister(c *net.Conn, resBuf []byte) error {
	fmt.Println("REGISTERR")
	conn := *c
	var p payload.RegisterLoginPayload
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
	if err != nil {
		return err
	}
	fmt.Printf("User create new account with username: %s and password: %s\n", p.Username, p.Password)
	res := business.Register(p.Username, p.Password)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleChangePassword(c *net.Conn, resBuf []byte) error {
	fmt.Println("CHANGE_PASSWORD")
	conn := *c
	var p payload.ChangePasswordPayload
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
	if err != nil {
		return err
	}
	fmt.Printf("User %s change password from %s to %s\n", p.Username, p.OldPassword, p.NewPassword)
	res := business.ChangePassword(p.Username, p.OldPassword, p.NewPassword)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}

func HandleChat(c *net.Conn, resBuf []byte) error {
	// fmt.Println("CHATTTTT")
	// var p payload.ChatPayload
	// err := utils.UnmarshalObject(&p, resBuf[4:])
	// if err != nil {
	// 	checkError(err)
	// }
	// fmt.Printf("%s send msg to %s with content: %s\n", p.From, p.To, p.Message)
	// conn.Write([]byte("ACKKKK"))
	return nil
}

func HandleFindUser(c *net.Conn, resBuf []byte) error {
	fmt.Println("FIND USER")
	conn := *c
	var p payload.UserPayload
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
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
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
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
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
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
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
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
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
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
	err := utils.UnmarshalObject(&p, resBuf[:len(resBuf)-1])
	if err != nil {
		return err
	}
	res := business.UserInfo(p.Username)
	resBytes := utils.MarshalObject(res)
	conn.Write(resBytes)
	return nil
}
