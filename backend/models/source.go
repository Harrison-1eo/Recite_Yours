package models

type Source struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"unique;not null"`
	UserID uint   `gorm:"not null"`
}

func IsSourceExist(name string, userID uint) bool {
	var source Source
	DB.Where("name = ? AND user_id = ?", name, userID).First(&source)
	return source.ID != 0
}
