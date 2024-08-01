package models

import "time"

type Word struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	UserID          uint
	Word            string
	Translation     string
	ExampleSentence string
	Note            string
	Tag             string
	SourceID        uint      `gorm:"not null"`
	Date            time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
