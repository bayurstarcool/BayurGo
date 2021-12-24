package database

import (
	"fmt"

	"github.com/bayurstarcool/BayurGo/app/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	dbUser := helpers.GetEnv("DB_USER")
	dbPass := helpers.GetEnv("DB_PASSWORD")
	dbName := helpers.GetEnv("DB_NAME")
	dbHost := helpers.GetEnv("DB_HOST")
	dbPort := helpers.GetEnv("DB_PORT")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	print("Database connected √\n\n")
	db.LogMode(true)
	return db
}
