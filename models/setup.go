package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/itms")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Hourly{})

	DB = db
}
