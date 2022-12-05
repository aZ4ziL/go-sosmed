package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = GetDB()
	db.AutoMigrate(
		&User{},
		&UserPost{},
		&UserPostComment{},
		&UserPostPhoto{},
	)
}

func GetDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/go-sosmed?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
