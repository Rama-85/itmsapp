package models

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//var hourly Hourly
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/itms"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Hourly{})

	DB = db

}
