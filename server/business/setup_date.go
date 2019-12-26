package business

import (
	message "final-project/message"
	"final-project/server/db/client"
	"final-project/server/db/entity"
	define "final-project/server/define"
	"time"

	"github.com/jinzhu/gorm"
)

const RFC3339FullDate = "2006-01-02"

func SetupDate(username string, date string) message.ReturnMessage {
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

	if !user.IsActive {
		return message.ReturnMessage{
			ReturnCode:    message.WrongPassword,
			ReturnMessage: message.GetMessageDecription(message.UserNotActive),
		}
	}

	t, _ := time.Parse(RFC3339FullDate, date)
	err = db.Model(&user).Update("birthday", t).Error
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
