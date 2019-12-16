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
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
