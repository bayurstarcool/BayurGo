package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func SetupDB() *gorm.DB {
	dbUser := getEnv("DB_USER")
	dbPass := getEnv("DB_PASSWORD")
	dbName := getEnv("DB_NAME")
	dbHost := getEnv("DB_HOST")
	dbPort := getEnv("DB_PORT")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	print("Database connected √\n\n")
	db.LogMode(true)
	return db
}
