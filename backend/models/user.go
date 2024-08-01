package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func IsUsernameExist(username string) bool {
	var user User
	DB.Where("username = ?", username).First(&user)
	return user.ID != 0
}
