package models

type EnChDict struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"ID"`
	Word     string `gorm:"index" json:"word"`
	Meaning  string `json:"meaning"`
	Priority int    `gorm:"default:0" json:"priority"`
}
