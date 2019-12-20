package cmd

import (
	"final-project/client/manager"
	"final-project/message"
	"final-project/server/constant"
	"final-project/utils"
	"fmt"
	"strings"
	_ "strings"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

var changePasswordCmd = &cobra.Command{
	Use:   "change_password",
	Short: "Change password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		fmt.Println("CHECK CHANGE PASSWORD")
		fmt.Print(">>password: ")
		pass, _ := gopass.GetPasswdMasked()
		passStr := strings.TrimRight(string(pass), "\n")
		fmt.Print(">> new password: ")
		newPass, _ := gopass.GetPasswdMasked()
		newPassStr := strings.TrimRight(string(newPass), "\n")
		clientService := manager.GetClientService()
		clientService.SendDataChangePassword(constant.Change_Password, args[0], passStr, newPassStr)
		conn := clientService.GetConnection()
		// utils.TellReadDone(&conn)
		var res message.ReturnMessage
		resData, _ := utils.ReadBytesResponse(&conn)
		err := utils.UnmarshalObject(&res, resData[:])
		if err != nil {
			fmt.Println("CANNOT UNMARSHAL")
			fmt.Println(err.Error())
			fmt.Println(string(resData[:]))
		}
		fmt.Printf("Return message: %d and %s\n", res.ReturnCode, res.ReturnMessage)
	},
}

func init() {
	rootCmd.AddCommand(changePasswordCmd)
}
