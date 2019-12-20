package cmd

import (
	"bufio"
	"final-project/client/manager"
	"final-project/message"
	"final-project/server/constant"
	"final-project/utils"
	"fmt"
	"os"
	"strings"
	_ "strings"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

var changePasswordCmd = &cobra.Command{
	Use:   "change_password",
	Short: "Change password",
	Long:  "User can change password by encrypting through flag encrypt",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("encrypt")
		passStr := ""
		newPassStr := ""
		if user != "" {
			passStr, newPassStr = getEncryptPassword()
		} else {
			fmt.Println(args[0])
			user = args[0]
			passStr, newPassStr = getUnencryptPassword()
		}
		fmt.Println("CHECK CHANGE PASSWORD")
<<<<<<< HEAD
		fmt.Print(">> password: ")
		pass, _ := reader.ReadString('\n')
		fmt.Print(">> new password: ")
		newPass, _ := reader.ReadString('\n')
=======
>>>>>>> daebdd76161a97160c4e76a43b7481e4dd49cefc
		clientService := manager.GetClientService()
		clientService.SendDataChangePassword(constant.Change_Password, user, passStr, newPassStr)
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
	},
}

func getEncryptPassword() (string, string) {
	fmt.Print(">>password: ")
	pass, _ := gopass.GetPasswdMasked()
	passStr := strings.TrimRight(string(pass), "\n")
	fmt.Print(">>new password: ")
	newPass, _ := gopass.GetPasswdMasked()
	newPassStr := strings.TrimRight(string(newPass), "\n")
	return passStr, newPassStr
}

func getUnencryptPassword() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimRight(pass, "\n")
	fmt.Print(">> new password: ")
	newPass, _ := reader.ReadString('\n')
	newPass = strings.TrimRight(newPass, "\n")
	return pass, newPass
}

func init() {
<<<<<<< HEAD
=======
	changePasswordCmd.PersistentFlags().StringP("encrypt", "e", "", "encrypt password before sending")
>>>>>>> daebdd76161a97160c4e76a43b7481e4dd49cefc
	rootCmd.AddCommand(changePasswordCmd)
}
