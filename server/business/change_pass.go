package business

import (
	message "final-project/message"
	"final-project/server/db/client"
	"final-project/server/db/entity"
	define "final-project/server/define"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(username string, oldPassword string, newPassword string) message.ReturnMessage {
	var user entity.User
	db := client.GetConnectionDB()
	err := db.Table(define.UserTable).Where("username = ?", username).First(&user).Error

	fmt.Println(user)

	if gorm.IsRecordNotFoundError(err) {
		return message.ReturnMessage{
			ReturnCode:    message.UsernameNotFound,
			ReturnMessage: message.GetMessageDecription(message.UsernameNotFound),
		}
	}
	if err != nil {
		return message.ReturnMessage{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return message.ReturnMessage{
			ReturnCode:    message.WrongPassword,
			ReturnMessage: message.GetMessageDecription(message.WrongPassword),
		}
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	err = db.Model(&user).Update("password", string(bytes)).Error
	if err != nil {
		return message.ReturnMessage{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
		}
	}
	return message.ReturnMessage{
		ReturnCode:    message.Success,
		ReturnMessage: message.GetMessageDecription(message.Success),
	}
}
