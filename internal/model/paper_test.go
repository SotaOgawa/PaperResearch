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

	db.AutoMigrate((&model.Paper{}))

	p := model.Paper{
		Title:      "Test Paper",
		Abstract:   "This is a test abstract.",
		Conference: "Test Conference",
		Year:       2023,
	}
	err = db.Create(&p).Error
	require.NoError(t, err)

	var result model.Paper
	err = db.First(&result, "title = ?", "Test Paper").Error
	require.NoError(t, err)
	require.Equal(t, "Test Paper", result.Title)
}
