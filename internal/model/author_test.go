package model_test

import (
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"paper-app-backend/internal/model"
	"testing"
)

func TestAuthor_SaveAndQuery(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	db.AutoMigrate((&model.AuthorObjectInDB{}))

	p := model.AuthorObjectInDB{
		Name: "Test Author",
		ID:   10,
	}
	err = db.Create(&p).Error
	require.NoError(t, err)

	var result model.AuthorObjectInDB
	err = db.First(&result, "ID = ?", 10).Error
	require.NoError(t, err)
	require.Equal(t, "Test Author", result.Name)
}
