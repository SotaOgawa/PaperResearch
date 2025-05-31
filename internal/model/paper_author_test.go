package model_test

import (
	"gorm.io/gorm"
	"testing"
	"gorm.io/driver/sqlite"
	"paper-app-backend/internal/model"
	"github.com/stretchr/testify/require"
)



func TestPaperAuthor_SaveAndQuery(t *testing.T){
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	db.AutoMigrate(&model.Author{}, &model.Paper{}, &model.PaperAuthorConnection{})

	p := model.Paper{
		Title: "Test Paper",
		Abstract: "This is a test abstract.",
		Conference: "Test Conference",
		Year: 2023,
		ID: 1,
	}

	a := model.Author{
		Name: "Test Author",
		ID: 10,
	}

	conn := model.PaperAuthorConnection{
		PaperID: p.ID,
		AuthorID: a.ID,
		AuthorOrder: 1,
	}

	err = db.Create(&p).Error
	require.NoError(t, err)
	err = db.Create(&a).Error
	require.NoError(t, err)
	err = db.Create(&conn).Error
	require.NoError(t, err)

	var result_paper model.Paper
	err = db.First(&result_paper, "Title = ?", "Test Paper").Error
	require.NoError(t, err)

	paper_ID := result_paper.ID

	var result_paper_author model.PaperAuthorConnection
	err = db.First(&result_paper_author, "paper_id = ?", paper_ID).Error
	require.NoError(t, err)

	var result model.Author
	err = db.First(&result, "id = ?", result_paper_author.AuthorID).Error
	require.NoError(t, err)

	require.Equal(t, "Test Paper", result_paper.Title)
	require.Equal(t, "Test Author", result.Name)
}