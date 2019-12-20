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

	"github.com/spf13/cobra"
)

var changePasswordCmd = &cobra.Command{
	Use:   "change_password",
	Short: "Change password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("CHECK CHANGE PASSWORD")
		fmt.Print(">>password: ")
		pass, _ := reader.ReadString('\n')
		pass = strings.TrimRight(pass, "\n")
		fmt.Print(">> new password: ")
		newPass, _ := reader.ReadString('\n')
		newPass = strings.TrimRight(newPass, "\n")
		clientService := manager.GetClientService()
		clientService.SendDataChangePassword(constant.Change_Password, args[0], pass, newPass)
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
