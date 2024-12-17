package database

import (
	"AUTH/src/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbError error

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")

	DB, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB!")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Token{})
	if err != nil{
		log.Fatal("Migration failed: ", err)
	} else{
		log.Println("Migration successful")
	}
}
