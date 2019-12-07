package cmd

import (
	"bufio"
	"final-project/client/manager"
	"fmt"
	"os"
	_ "strings"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with user's account",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("CHECK LOGINNNN")
		fmt.Println(">>Password: ")
		pass, _ := reader.ReadString('\n')
		clientService := manager.GetClientService()
		clientService.SendData(args[0] + "\n" + pass)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}