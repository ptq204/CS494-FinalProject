package cmd

import (
	"final-project/client/manager"
	"final-project/constant"
	"final-project/message"
	"final-project/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var setupUserInfoCmd = &cobra.Command{
	Use:   "setup",
	Short: "Change user's info",
	Long:  "Change user's info",
	Run: func(cmd *cobra.Command, args []string) {
		setupUserActions := []int{
			constant.SetupName,
			constant.SetupDate,
			constant.SetupNote,
		}

		if len(args) != 1 {
			fmt.Println("Wrong command's format")
			return
		}

		setupUserFlags := []string{
			"name", "date", "note",
		}

		clientService := manager.GetClientService()
		var username string
		var newInfo string

		for idx, f := range setupUserFlags {
			username, _ = cmd.Flags().GetString(f)
			if username != "" {
				newInfo = args[0]
				clientService.SetupInfo(setupUserActions[idx], username, newInfo)
				break
			}
		}

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
	setupUserInfoCmd.PersistentFlags().StringP("name", "n", "", "Change user's name")
	setupUserInfoCmd.PersistentFlags().StringP("date", "d", "", "Change user's birthday")
	setupUserInfoCmd.PersistentFlags().StringP("note", "o", "", "Change user's note")
	rootCmd.AddCommand(setupUserInfoCmd)
}
