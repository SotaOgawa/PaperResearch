package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"paper-app-backend/internal/db"
	"paper-app-backend/internal/model"
)

func GetPapers(c *gin.Context) {
	var papers []model.Paper
	var paperQuery PaperQuery

	err := c.ShouldBindQuery(&paperQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	queryFiltered := paperQuery.Apply(db.DB.Model(&model.Paper{}))

	err = queryFiltered.Find(&papers).Error; 

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, papers)
}
