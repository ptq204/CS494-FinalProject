package cmd

import (
	"bufio"
	"final-project/client/manager"
	"final-project/server/constant"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	_ "strings"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with user's account",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("CHECK LOGINNNN")
		fmt.Print(">>Password: ")
		pass, _ := reader.ReadString('\n')
		clientService := manager.GetClientService()
		clientService.SendDataRegisterLogin(constant.Login, args[0], pass)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
