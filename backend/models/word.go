package models

import "time"

type Word struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"ID"`
	UserID          uint      `gorm:"not null" json:"userID"`
	Word            string    `gorm:"not null" json:"word"`
	Translation     string    `json:"translation"`
	ExampleSentence string    `json:"exampleSentence"`
	Note            string    `json:"note"`
	Tag             string    `json:"tag"`
	SourceID        uint      `gorm:"not null" json:"sourceID"`
	Date            time.Time `gorm:"not null" json:"date"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
