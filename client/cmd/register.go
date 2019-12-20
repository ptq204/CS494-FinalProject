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

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register new account",
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
		fmt.Println("CHECK REGISTER")
		clientService := manager.GetClientService()
		clientService.SendDataRegisterLogin(constant.Register, user, passStr)
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

func getRegisterEncryptPassword() string {
	fmt.Print(">>password: ")
	pass, _ := gopass.GetPasswdMasked()
	passStr := strings.TrimRight(string(pass), "\n")
	return passStr
}

func getRegisterUnencryptPassword() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimRight(pass, "\n")
	return pass
}
func init() {
	registerCmd.PersistentFlags().StringP("encrypt", "e", "", "encrypt password before sending")
	rootCmd.AddCommand(registerCmd)
}
