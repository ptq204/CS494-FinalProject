package entity

type User struct {
	ID       int32  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type File struct {
	ID       int32  `gorm:"column:id"`
	Filename string `gorm:"column:filename"`
	Data
}
