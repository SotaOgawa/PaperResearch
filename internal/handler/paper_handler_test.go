package handler_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/require"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "paper-app-backend/internal/handler"
    "paper-app-backend/internal/model"
)

func setupTestRouter(db *gorm.DB) *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/api/papers", func(c *gin.Context) {
        var papers []model.Paper
        var paperQuery handler.PaperQuery
        if err := c.ShouldBindQuery(&paperQuery); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        filtered := paperQuery.Apply(db.Model(&model.Paper{}))
        if err := filtered.Find(&papers).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, papers)
    })
    return r
}

func TestGetPapers_WithQuery(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&model.Paper{})
    db.Create(&model.Paper{Title: "Transformer", Conference: "ICLR", Year: 2023})

    router := setupTestRouter(db)

    req, _ := http.NewRequest("GET", "/api/papers?title=Transformer", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    require.Equal(t, 200, w.Code)
    require.Contains(t, w.Body.String(), "Transformer")
}

func TestGetPapers_EmptyResult(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&model.Paper{})

    router := setupTestRouter(db)

    req, _ := http.NewRequest("GET", "/api/papers?title=Nonexistent", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    require.Equal(t, 200, w.Code)
    require.NotContains(t, w.Body.String(), "Nonexistent")
}
