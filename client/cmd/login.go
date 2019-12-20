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

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with user's account",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("encrypt")
		passStr := ""
		if user != "" {
			passStr = getLoginEncryptPassword()
		} else {
			fmt.Println(args[0])
			user = args[0]
			passStr = getLoginUnencryptPassword()
		}
		fmt.Println("CHECK LOGINNNN")
		clientService := manager.GetClientService()
		clientService.SendDataRegisterLogin(constant.Login, user, passStr)
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

func getLoginEncryptPassword() string {
	fmt.Print(">>Password: ")
	pass, _ := gopass.GetPasswdMasked()
	passStr := strings.TrimRight(string(pass), "\n")
	return passStr
}

func getLoginUnencryptPassword() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimRight(pass, "\n")
	return pass
}

func init() {
	loginCmd.PersistentFlags().StringP("encrypt", "e", "", "encrypt password before sending")
	rootCmd.AddCommand(loginCmd)
}
