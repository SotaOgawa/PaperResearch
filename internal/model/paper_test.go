package model_test

import (
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"paper-app-backend/internal/model"
	"testing"
)

func TestPaper_SaveAndQuery(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	db.AutoMigrate((&model.PaperObjectInDB{}))

	p := model.PaperObjectInDB{
		Title:      "Test Paper",
		Abstract:   "This is a test abstract.",
		Authors:    "John Doe, Jane Smith",
		Conference: "Test Conference",
		Year:       2023,
	}
	err = db.Create(&p).Error
	require.NoError(t, err)

	var result model.PaperObjectInDB
	err = db.First(&result, "title = ?", "Test Paper").Error
	require.NoError(t, err)
	require.Equal(t, "Test Paper", result.Title)
	require.Equal(t, "This is a test abstract.", result.Abstract)
	require.Equal(t, "John Doe, Jane Smith", result.Authors)
	require.Equal(t, "Test Conference", result.Conference)
	require.Equal(t, 2023, result.Year)
	require.NotZero(t, result.ID) // ID should be set after creation
}
