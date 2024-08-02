package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open("dictionarymaster:dictionarypasswd@tcp(127.0.0.1:3306)/dictionarydb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	// 自动迁移模式
	err = DB.AutoMigrate(
		&User{},
		&Source{},
		&Word{},
		&EnChDict{},
	)
	if err != nil {
		panic("failed to migrate database" + err.Error())
	}
}
