package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func IsUsernameExist(username string) bool {
	var user User
	DB.Where("username = ?", username).First(&user)
	return user.ID != 0
}
