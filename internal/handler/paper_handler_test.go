package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"bytes"
	"paper-app-backend/internal/handler"
	"paper-app-backend/internal/model"
	"strconv"
)

func setupGetRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/papers", func(c *gin.Context) {
		handler.GetPapersWithDB(c, db)
	})
	return r
}

func setupPostRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/papers", func(c *gin.Context) {
		handler.CreatePaperWithDB(c, db)
	})
	return r
}

func setupPutRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/papers/:id", func(c *gin.Context) {
		handler.UpdatePaperWithDB(c, db)
	})
	return r
}

func setupDeleteRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/api/papers/:id", func(c *gin.Context) {
		handler.DeletePaperWithDB(c, db)
	})
	return r
}

func TestGetPapers_WithQuery(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})
	db.Create(&model.PaperObjectInDB{Title: "Transformer", Conference: "ICLR", Year: 2023})

	router := setupGetRouter(db)

	req, _ := http.NewRequest("GET", "/api/papers?title=Transformer", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
	require.Contains(t, w.Body.String(), "Transformer")
}

func TestGetPapers_EmptyResult(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})

	router := setupGetRouter(db)

	req, _ := http.NewRequest("GET", "/api/papers?title=Nonexistent", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
	require.NotContains(t, w.Body.String(), "Nonexistent")
}

func TestCreatePaper_Success(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})

	router := setupPostRouter(db)

	body := `{
		"title": "Test Paper",
		"conference": "TestConf",
		"year": 2024,
		"authors": "John Doe, Jane Smith",
		"abstract": "This is a test",
		"url": "https://example.com",
		"citation_count": 5,
		"bibtex": "@article{...}",
		"pdf_url": "https://example.com/test.pdf"
	}`

	req, _ := http.NewRequest("POST", "/api/papers", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var createdPaper model.PaperObjectInDB
	err := db.First(&createdPaper, "title = ?", "Test Paper").Error
	require.NoError(t, err)
	require.Equal(t, "Test Paper", createdPaper.Title)
	require.Equal(t, "TestConf", createdPaper.Conference)
	require.Equal(t, 2024, createdPaper.Year)
	require.Equal(t, "John Doe, Jane Smith", createdPaper.Authors)
	require.Equal(t, "This is a test", createdPaper.Abstract)
	require.Equal(t, "https://example.com", createdPaper.URL)
	require.Equal(t, 5, createdPaper.CitationCount)
	require.Equal(t, "@article{...}", createdPaper.Bibtex)
	require.Equal(t, "https://example.com/test.pdf", createdPaper.PDFURL)
	require.NotZero(t, createdPaper.ID) // ID should be set after creation
}

func TestCreatePaper_WithID_ShouldFail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})

	router := setupPostRouter(db)

	// IDを指定している
	body := `{
		"id": 1,
		"title": "Should Fail",
		"year": 2024
	}`

	req, _ := http.NewRequest("POST", "/api/papers", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Contains(t, w.Body.String(), "ID should not be provided")
}

func TestCreatePaper_InvalidJSON_ShouldFail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})

	router := setupPostRouter(db)

	body := `{ "title": "Incomplete JSON", ` // ← JSON壊れてる

	req, _ := http.NewRequest("POST", "/api/papers", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Contains(t, w.Body.String(), "invalid input data")
}

func TestUpdatePaper_Success(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})

	// まずはPaperを作成
	paper := model.PaperObjectInDB{
		Title:      "Old Title",
		Conference: "Old Conference",
		Year:       2023,
	}
	db.Create(&paper)

	router := setupPutRouter(db)

	body := `{
		"id": ` + strconv.Itoa(paper.ID) + `,
		"title": "Updated Title",
		"conference": "Updated Conference",
		"year": 2024
	}`

	req, _ := http.NewRequest("PUT", "/api/papers/"+strconv.Itoa(paper.ID), bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "Updated Title")
}

func TestDeletePaper_Success(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.PaperObjectInDB{})

	// まずはPaperを作成
	paper := model.PaperObjectInDB{
		Title:      "Paper to Delete",
		Conference: "Conference",
		Year:       2023,
	}
	db.Create(&paper)

	router := setupDeleteRouter(db)

	req, _ := http.NewRequest("DELETE", "/api/papers/"+strconv.Itoa(paper.ID), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)

	var deletedPaper model.PaperObjectInDB
	err := db.First(&deletedPaper, paper.ID).Error
	require.Error(t, err) // Paperが削除されているはず
}
