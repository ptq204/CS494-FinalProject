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

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with user's account",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("CHECK LOGINNNN")
		fmt.Print(">>Password: ")
		pass, _ := reader.ReadString('\n')
		pass = strings.TrimRight(pass, "\n")
		clientService := manager.GetClientService()
		clientService.SendDataRegisterLogin(constant.Login, args[0], pass)
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
	rootCmd.AddCommand(loginCmd)
}
