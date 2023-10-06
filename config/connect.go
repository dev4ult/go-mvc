package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "praktikum/models"
)

type Connection struct {
	db_user 	string
	db_pass 	string
	db_host 	string
	db_port 	int
	db_name 	string
	SERVER_PORT	int
}


func InitDB(conn Connection) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conn.db_user, conn.db_pass, conn.db_host, conn.db_port, conn.db_name)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	initMigrate(db)

	return db
}

func LoadConfig() Connection {
	godotenv.Load(".env")

	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT")) 
	SERVER_PORT, _ := strconv.Atoi(os.Getenv("SERVER_PORT")) 

	return Connection{
		db_user : os.Getenv("DB_USER"),
		db_pass : os.Getenv("DB_PASS"),
		db_host : os.Getenv("DB_HOST"),
		db_port : DB_PORT,
		db_name : os.Getenv("DB_NAME"),
		SERVER_PORT : SERVER_PORT,
	}
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, model.Book{}, &model.Blog{})
}