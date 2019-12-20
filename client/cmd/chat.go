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

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with one or multiple users",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("CHECK CHAT")
		for {
			fmt.Print(">>Me: ")
			msg, _ := reader.ReadString('\n')
			clientService := manager.GetClientService()
			clientService.SendDataChat(constant.Chat, "quyen", args[0], msg)
			res, err := clientService.ReadData()
			if err != nil {
				break
			}
			fmt.Println("Response: " + res)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
