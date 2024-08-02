package models

type Source struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"ID"`
	Name   string `gorm:"unique;not null" json:"name"`
	UserID uint   `gorm:"not null" json:"userID"`
}

func IsSourceExist(name string, userID uint) bool {
	var source Source
	DB.Where("name = ? AND user_id = ?", name, userID).First(&source)
	return source.ID != 0
}
