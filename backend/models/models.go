package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
}

type Source struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type Word struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	Word            string
	Translation     string
	ExampleSentence string
	Note            string
	SourceID        uint
	Date            time.Time
	CreatedAt       time.Time
}
