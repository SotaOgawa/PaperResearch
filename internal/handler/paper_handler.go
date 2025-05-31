package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"paper-app-backend/internal/model"
	"paper-app-backend/internal/db"
)

func GetPapers(c *gin.Context) {
	GetPapersWithDB(c, db.DB)
}

func GetPapersWithDB(c *gin.Context, db *gorm.DB) {
	var papers []model.Paper
	var paperQuery PaperSearchQuery

	err := c.ShouldBindQuery(&paperQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	queryFiltered := paperQuery.Apply(db)

	err = queryFiltered.Find(&papers).Error; 

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, papers)
}

func CreatePaper(c *gin.Context) {
	CreatePaperWithDB(c, db.DB)
}

func CreatePaperWithDB(c *gin.Context, db *gorm.DB) {
	var newPaper model.Paper

	err := c.ShouldBindJSON(&newPaper); 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	if newPaper.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should not be provided for new papers"})
		return
	}

	if err := db.Create(&newPaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPaper)
}

func UpdatePaper(c *gin.Context) {
	UpdatePaperWithDB(c, db.DB)
}

func UpdatePaperWithDB(c *gin.Context, db *gorm.DB) {
	var updatedPaper model.Paper

	err := c.ShouldBindJSON(&updatedPaper); 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	if updatedPaper.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided for updates"})
		return
	}

	var existing model.Paper
	if err := db.First(&existing, updatedPaper.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paper not found"})
		return
	}

	if err := db.Model(&existing).Updates(updatedPaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPaper)
}