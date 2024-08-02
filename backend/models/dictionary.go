package models

type EnChDict struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Word     string `gorm:"index"`
	Meaning  string
	Priority int `gorm:"default:0"`
}
