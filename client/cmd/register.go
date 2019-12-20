package cmd

import (
	"bufio"
	"final-project/client/manager"
	"final-project/message"
	"final-project/server/constant"
	"final-project/utils"
	"fmt"
	"os"
	_ "strings"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register new account",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("CHECK REGISTER")
		fmt.Print(">>Password: ")
		pass, _ := reader.ReadString('\n')
		clientService := manager.GetClientService()
		clientService.SendDataRegisterLogin(constant.Register, args[0], pass)
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

func init() {
	rootCmd.AddCommand(registerCmd)
}
