package db

import (
	"fmt"
	"log"
	"os"
	"paper-app-backend/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env.local")

	dsn := os.Getenv("POSTGRES_URL")
	fmt.Println("Connecting to database with DSN:", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// マイグレーション
	err = DB.AutoMigrate(&model.PaperObjectInDB{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database initialized successfully")
}
