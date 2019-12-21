package cmd

import (
	"final-project/client/manager"
	"final-project/message"
	"final-project/constant"
	"final-project/utils"
	"fmt"
	_ "strings"

	"github.com/spf13/cobra"
)

var checkUserCmd = &cobra.Command{
	Use:   "check_user",
	Short: "Check user info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("Wrong command's format")
			return
		}

		clientService := manager.GetClientService()

		checkUserFlags := []string{
			"find", "online", "show_date", "show_note", "show_name", "show",
		}
		checkUserActions := []int{
			constant.FindUser,
			constant.UserOnline,
			constant.UserBirthday,
			constant.UserNote,
			constant.UserName,
			constant.UserInfo,
		}

		checkShowInfo := false
		var username string

		for idx, f := range checkUserFlags {
			username, _ = cmd.Flags().GetString(f)
			if username != "" {
				clientService.CheckUser(checkUserActions[idx], username)
				if checkUserActions[idx] == constant.UserInfo {
					checkShowInfo = true
				}
				break
			}
		}

		if !checkShowInfo {
			var res message.CheckUserResponse
			conn := clientService.GetConnection()
			resData, _ := utils.ReadBytesResponse(&conn)
			err := utils.UnmarshalObject(&res, resData[:])
			if err != nil {
				fmt.Println("CANNOT UNMARSHAL")
				fmt.Println(err.Error())
				fmt.Println(string(resData[:]))
			}
			fmt.Printf("Info: %s with %d and %s\n", res.Information, res.ReturnMessage.ReturnCode, res.ReturnMessage.ReturnMessage)
		} else {
			var res message.UserResponseInfo
			conn := clientService.GetConnection()
			resData, _ := utils.ReadBytesResponse(&conn)
			err := utils.UnmarshalObject(&res, resData[:])
			if err != nil {
				fmt.Println("CANNOT UNMARSHAL")
				fmt.Println(err.Error())
				fmt.Println(string(resData[:]))
			}
			fmt.Print("User info: ", res.User, " with ", res.ReturnCode, " and ", res.ReturnMessage, "\n")
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
