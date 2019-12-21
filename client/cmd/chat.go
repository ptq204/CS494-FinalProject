package cmd

import (
	"bufio"
	"final-project/client/manager"
	"final-project/constant"
	"final-project/message"
	"final-project/security"
	"final-project/utils"
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
		token, err := utils.ReadLocalValueInFile("token")

		if err != nil || token == "" {
			fmt.Println("Please login before joining chat room")
		}

		payLoad, err := security.ParseToken(token)
		if err != nil {
			fmt.Println("Cannot authenticate. Please login again")
			return
		}

		username := payLoad["username"].(string)
		clientService := manager.GetClientService()
		conn := clientService.GetConnection()

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

			clientService.SendDataChat(constant.Chat, username, args[0], msg)
			var res message.ReturnMessageChat
			resData, _ := utils.ReadBytesResponse(&conn)
			err := utils.UnmarshalObject(&res, resData[:])

			if err != nil {
				break
			}

			fmt.Printf("%s : %s\n", res.From, res.Message)
		}
	},
}

func init() {
	chatCmd.PersistentFlags().StringP("encrypt", "e", "", "encrypt message")
	rootCmd.AddCommand(chatCmd)
}
