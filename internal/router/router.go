package router

import (
	"github.com/gin-gonic/gin"
	"paper-app-backend/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/papers", handler.GetPapers)            // 一覧取得（検索あり）
		// api.POST("/papers", handler.CreatePaper)         // 新規追加
		// api.PUT("/papers/:id", handler.UpdatePaper)      // 更新
		// api.DELETE("/papers/:id", handler.DeletePaper)   // 削除
	}
}