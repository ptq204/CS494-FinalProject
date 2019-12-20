package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID  `gorm:"primary_key;column:id" json:"-"`
	Name     string     `gorm:"column:name" json:"name"`
	Username string     `gorm:"unique_index;column:username" json:"username"`
	Password string     `gorm:"column:password" json:"-"`
	Note     string     `gorm:"column:note" json:"note"`
	Birthday *time.Time `gorm:"column:birthday" json:"birthday"`
	IsActive bool       `gorm:"column:is_active" json:"is_active"`
}
type Channel struct {
	ID uuid.UUID `gorm:"column:id"`
}
type File struct {
	ID       uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	Filename string    `gorm:"unique;column:name"`
	Content  []byte    `sql:"type:blob" gorm:"column:content"`
	UserID   uuid.UUID `sql:"REFERENCES user(id)" gorm:"column:user_id"`
}

type Message struct {
	FromUser  uuid.UUID `sql:"references user(id)" gorm:"column:from_user"`
	ToUser    uuid.UUID `sql:"references user(id)" gorm:"column:to_user"`
	ChannelID uuid.UUID `sql:"references channel(id)" gorm:"column:channel_id"`
	Message   string    `gorm:"column:message"`
}

type UserChannel struct {
	UserID    uuid.UUID `sql:"references user(id)" gorm:"column:user_id"`
	ChannelID Channel   `sql:"references channel(id)" gorm:"column:channel_id"`
}
