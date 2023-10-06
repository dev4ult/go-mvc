package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "praktikum/models"
)


func InitDB() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/basic_connect_golang?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	initMigrate(db)

	return db
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, model.Book{}, &model.Blog{})
}