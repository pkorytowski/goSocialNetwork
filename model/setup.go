package model

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User{}, &LoginData{}, &Post{}, &Like{})
	if err != nil {
		return
	}

	DB = database
}
