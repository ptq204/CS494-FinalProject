package business

import (
	"final-project/server/db/client"
	"final-project/server/db/entity"
	define "final-project/server/define"
	message "final-project/message"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Signin(username string, password string) message.ReturnMessage{
	// signin
	var user entity.User
	db := client.GetConnectionDB()
	err := db.Table(define.UserTable).Where("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return message.ReturnMessage {
			ReturnCode: message.UsernameNotFound,
			ReturnMessage: message.GetMessageDecription(message.UsernameNotFound),
		}
	}
	if err != nil {
		return message.ReturnMessage{
			ReturnCode: message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return message.ReturnMessage{
			ReturnCode: message.WrongPassword,
			ReturnMessage: message.GetMessageDecription(message.WrongPassword),
		}
	}
	return message.ReturnMessage{
		ReturnCode: message.Success,
		ReturnMessage: message.GetMessageDecription(message.Success),
	}
}
