package model_test

import (
	"gorm.io/gorm"
	"testing"
	"gorm.io/driver/sqlite"
	"paper-app-backend/internal/model"
	"github.com/stretchr/testify/require"
)



func TestAuthor_SaveAndQuery(t *testing.T){
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	db.AutoMigrate((&model.Author{}))

	p := model.Author{
		Name: "Test Author",
		ID: 10,
	}
	err = db.Create(&p).Error
	require.NoError(t, err)

	var result model.Author
	err = db.First(&result, "ID = ?", 10).Error
	require.NoError(t, err)
	require.Equal(t, "Test Author", result.Name)
}