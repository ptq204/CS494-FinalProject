package business

import (
	"final-project/constant"
	"final-project/decrypt"
	"final-project/message"
	"final-project/server/db/client"
	"final-project/server/db/entity"
	"final-project/server/define"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(username string, password string, encrypt int32) message.ReturnMessage {
	if encrypt == 1 {
		username = decrypt.Decrypt(constant.PASSPHRASE, username)
		password = decrypt.Decrypt(constant.PASSPHRASE, password)
	}
	db := client.GetConnectionDB()
	var user entity.User
	err := db.Table(define.UserTable).Where("username=?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return message.ReturnMessage{
			ReturnCode:    message.Unknown,
			ReturnMessage: message.GetMessageDecription(message.Unknown),
		}
	}
	if err == nil {
		return message.ReturnMessage{
			ReturnCode:    message.UsernameExist,
			ReturnMessage: message.GetMessageDecription(message.UsernameExist),
		}
	} else {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		user := entity.User{Username: username, Password: string(bytes), ID: uuid.NewV4(), IsActive: false}
		err = db.Table("user").Create(&user).Error
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
}
