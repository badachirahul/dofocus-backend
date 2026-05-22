package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/badachirahul/dofocus-backend/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	DB = database

	err = DB.AutoMigrate(
		&models.User{},
		&models.OTPVerification{},
		&models.Task{},
		&models.FocusSession{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	log.Println("Database connected successfully")
}
