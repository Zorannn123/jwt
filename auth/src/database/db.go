package database

import (
	"AUTH/src/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbError error

func Connect() {
	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=123 sslmode=disable"
	DB, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB!")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed")
}
