package business

import (
	"final-project/message"
	"final-project/server/db/client"
	"final-project/server/db/entity"
	define "final-project/server/define"

	"github.com/jinzhu/gorm"
)

func Exit(username string) message.ReturnMessage {
	var user entity.User
	db := client.GetConnectionDB()
	err := db.Table(define.UserTable).Where("username = ?", username).First(&user).Error

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
	err = db.Model(&user).Update("is_active", false).Error
	if err != nil {
		return message.ReturnMessage{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.CannotSetupUserInfo),
		}
	}
	return message.ReturnMessage{
		ReturnCode:    message.Success,
		ReturnMessage: message.GetMessageDecription(message.Success),
	}
}
