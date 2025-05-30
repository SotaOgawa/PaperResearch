package db

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"paper-app-backend/internal/model"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("papers.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// マイグレーション
	err = DB.AutoMigrate(&model.Paper{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database initialized successfully")
}