package business

import (
	"final-project/message"
	"final-project/server/db/client"
	"final-project/server/db/entity"
	"final-project/server/define"

	"github.com/jinzhu/gorm"
)

func OnlineUser(username string) message.CheckUserResponse {
	db := client.GetConnectionDB()
	var user entity.User
	err := db.Table(define.UserTable).Where("username=?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return message.CheckUserResponse{
			Information: "",
			ReturnMessage: message.ReturnMessage{
				ReturnCode:    message.UsernameNotFound,
				ReturnMessage: message.GetMessageDecription(message.UsernameNotFound),
			},
		}
	}
	if err != nil {
		return message.CheckUserResponse{
			Information: "",
			ReturnMessage: message.ReturnMessage{
				ReturnCode:    message.Unknown,
				ReturnMessage: message.GetMessageDecription(message.Unknown),
			},
		}
	} else {
		if user.IsActive == false {
			return message.CheckUserResponse{
				Information: "User is not online",
				ReturnMessage: message.ReturnMessage{
					ReturnCode:    message.Success,
					ReturnMessage: message.GetMessageDecription(message.Success),
				},
			}
		} else {
			return message.CheckUserResponse{
				Information: "User is online",
				ReturnMessage: message.ReturnMessage{
					ReturnCode:    message.Success,
					ReturnMessage: message.GetMessageDecription(message.Success),
				},
			}
		}
	}
}
