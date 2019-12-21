package business

import (
	message "final-project/message"
	"final-project/server/db/client"
	"final-project/server/db/entity"
	define "final-project/server/define"

	security "final-project/security"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Signin(username string, password string) message.ReturnMessageLogin {
	// signin
	var user entity.User
	db := client.GetConnectionDB()
	err := db.Table(define.UserTable).Where("username = ?", username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return message.ReturnMessageLogin{
			ReturnCode:    message.UsernameNotFound,
			ReturnMessage: message.GetMessageDecription(message.UsernameNotFound),
			Token:         "",
		}
	}
	if err != nil {
		return message.ReturnMessageLogin{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
			Token:         "",
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return message.ReturnMessageLogin{
			ReturnCode:    message.WrongPassword,
			ReturnMessage: message.GetMessageDecription(message.WrongPassword),
			Token:         "",
		}
	}
	err = db.Model(&user).Update("is_active", true).Error
	if err != nil {
		return message.ReturnMessageLogin{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
			Token:         "",
		}
	}
	tokenPayload := map[string]string{
		"username": username,
	}
	tokenStr, err := security.GenerateToken(tokenPayload)

	if err != nil {
		return message.ReturnMessageLogin{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
			Token:         "",
		}
	}
	return message.ReturnMessageLogin{
		ReturnCode:    message.Success,
		ReturnMessage: message.GetMessageDecription(message.Success),
		Token:         tokenStr,
	}
}
