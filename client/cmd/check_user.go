package cmd

import (
	"final-project/client/manager"
	"final-project/message"
	"final-project/server/constant"
	"final-project/utils"
	"fmt"
	_ "strings"

	"github.com/spf13/cobra"
)

var checkUserCmd = &cobra.Command{
	Use:   "check_user",
	Short: "Check user info",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("find")
		clientService := manager.GetClientService()
		if username != "" {
			fmt.Println("FIND IF USER IS EXISTED")
			clientService.CheckUser(constant.FindUser, username)
		} else {
			username, _ = cmd.Flags().GetString("online")
			if username != "" {
				fmt.Println("CHECK IF USER ONLINE")
				clientService.CheckUser(constant.UserOnline, username)
			} else {
				username, _ = cmd.Flags().GetString("show_date")
				if username != "" {
					fmt.Println("SHOW USER'S BIRTHDAY")
					clientService.CheckUser(constant.UserBirthday, username)
				} else {
					username, _ = cmd.Flags().GetString("show_note")
					if username != "" {
						fmt.Println("SHOW USER'S NOTE")
						clientService.CheckUser(constant.UserNote, username)
					} else {
						username, _ = cmd.Flags().GetString("show_name")
						if username != "" {
							fmt.Println("SHOW USER'S NAME")
							clientService.CheckUser(constant.UserName, username)
						}
					}
				}
			}
		}
		conn := clientService.GetConnection()
		if username != "" {
			var res message.CheckUserResponse
			resData, _ := utils.ReadBytesResponse(&conn)
			err := utils.UnmarshalObject(&res, resData[:])
			if err != nil {
				fmt.Println("CANNOT UNMARSHAL")
				fmt.Println(err.Error())
				fmt.Println(string(resData[:]))
			}
			fmt.Printf("Info: %s with %d and %s\n", res.Information, res.ReturnMessage.ReturnCode, res.ReturnMessage.ReturnMessage)
		} else {
			username, _ = cmd.Flags().GetString("show")
			if username != "" {
				fmt.Println("SHOW USER'S INFO")
				clientService.CheckUser(constant.UserInfo, username)
				var res message.UserResponseInfo
				resData, _ := utils.ReadBytesResponse(&conn)
				err := utils.UnmarshalObject(&res, resData[:])
				if err != nil {
					fmt.Println("CANNOT UNMARSHAL")
					fmt.Println(err.Error())
					fmt.Println(string(resData[:]))
				}
				fmt.Print("User info: ", res.User, " with ", res.ReturnCode, " and ", res.ReturnMessage, "\n")
			}
		}
	},
}

func init() {
	checkUserCmd.PersistentFlags().StringP("encrypt", "e", "", "encrypt password before sending")
	checkUserCmd.PersistentFlags().StringP("find", "f", "", "find if user is existed")
	checkUserCmd.PersistentFlags().StringP("online", "o", "", "check if user online")
	checkUserCmd.PersistentFlags().StringP("show_date", "d", "", "show user's date")
	checkUserCmd.PersistentFlags().StringP("show_name", "n", "", "show user's name")
	checkUserCmd.PersistentFlags().StringP("show_note", "t", "", "show user's info")
	checkUserCmd.PersistentFlags().StringP("show", "s", "", "show user's note")
	rootCmd.AddCommand(checkUserCmd)
}
