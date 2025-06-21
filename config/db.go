package config

import (
	"fmt"
	"log"
	"os"
	"ticketing-system/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) // at top with other imports

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	DB = db
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.User{}, &models.Ticket{})
	DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Ticket{})

	log.Println("✅ Tables migrated")
	log.Println("✅ Database connected")
}
