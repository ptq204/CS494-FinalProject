package cmd

import (
	"bufio"
	"final-project/client/manager"
	"final-project/constant"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	_ "strings"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with one or multiple users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CHECK CHAT")
		clientService := manager.GetClientService()
		currUserName := clientService.GetCurrUserName()

		if currUserName == "" {
			fmt.Println("You have to login before joining chat room")
			return
		}

		var msg string
		encryptCheck, _ := cmd.Flags().GetString("encrypt")
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print(">> Me: ")
			msg, _ = reader.ReadString('\n')

			if encryptCheck != "" {
				// Put encrypt function here
				msg = "ENCRYPT MESSAGE HERE"
			}

			clientService.SendDataChat(constant.Chat, currUserName, args[0], msg)
			res, err := clientService.ReadData()
			if err != nil {
				break
			}

			fmt.Println("Response: " + res)
		}
	},
}

func init() {
	chatCmd.PersistentFlags().StringP("encrypt", "e", "", "encrypt message")
	rootCmd.AddCommand(chatCmd)
}
