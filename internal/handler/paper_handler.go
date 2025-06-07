package handler

import (
	"fmt"
	"net/http"
	"paper-app-backend/internal/db"
	"paper-app-backend/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPapers(c *gin.Context) {
	GetPapersWithDB(c, db.DB)
}

func GetPapersWithDB(c *gin.Context, db *gorm.DB) {
	var papers []model.PaperObjectInDB
	var paperQuery PaperSearchQuery

	err := c.ShouldBindQuery(&paperQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queryFiltered := paperQuery.Apply(db)

	err = queryFiltered.Find(&papers).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"papers": papers,
		"count":  len(papers)})
}

func CreatePaper(c *gin.Context) {
	CreatePaperWithDB(c, db.DB)
}

func CreatePaperWithDB(c *gin.Context, db *gorm.DB) {
	var newPaper model.PaperObjectInDB

	err := c.ShouldBindJSON(&newPaper)

	if err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newPaper.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID should not be provided for new papers"})
		return
	}

	// Check if the paper with the same title already exists
	err = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "title"}},
		DoUpdates: clause.AssignmentColumns([]string{"abstract", "citation_count", "updated_at", "bibtex", "pdf_url", "url"}),
	}).Create(&newPaper).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"paper": newPaper})
}

func UpdatePaper(c *gin.Context) {
	UpdatePaperWithDB(c, db.DB)
}

func UpdatePaperWithDB(c *gin.Context, db *gorm.DB) {
	var updatedPaper model.PaperObjectInDB

	err := c.ShouldBindJSON(&updatedPaper)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorr": "invalid input data"})
		return
	}

	if updatedPaper.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be provided for updates"})
		return
	}

	var existing model.PaperObjectInDB
	if err := db.First(&existing, updatedPaper.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paper not found"})
		return
	}

	if err := db.Model(&existing).Updates(updatedPaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"paper": updatedPaper})
}

func DeletePaper(c *gin.Context) {
	DeletePaperWithDB(c, db.DB)
}

func DeletePaperWithDB(c *gin.Context, db *gorm.DB) {
	var paper model.PaperObjectInDB
	var err error

	paper.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid paper ID"})
		return
	}

	if err := db.First(&paper).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paper not found"})
		return
	}

	if err := db.Delete(&paper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
