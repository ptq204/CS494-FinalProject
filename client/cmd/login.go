package cmd

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"final-project/client/manager"
	"final-project/message"
	"final-project/server/constant"
	"final-project/utils"
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
		fmt.Print(">>Password: ")
		pass, _ := reader.ReadString('\n')
		clientService := manager.GetClientService()
		clientService.SendDataRegisterLogin(constant.Login, args[0], pass)
		conn := clientService.GetConnection()
		var res message.ReturnMessage
		resData, err := utils.ReadBytesData(&conn)
		tmpBuff := bytes.NewBuffer(resData)
		utils.CheckError(err)
		d := gob.NewDecoder(tmpBuff)
		d.Decode(&res)
		fmt.Printf("Return message: %d and %s\n", res.ReturnCode, res.ReturnMessage)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
