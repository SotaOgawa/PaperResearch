package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"paper-app-backend/internal/db"
	"paper-app-backend/internal/router"
)

func main() {
	// データベースの初期化
	db.InitDB()

	// ルーターのセットアップ
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},        // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // 許可するHTTPメソッド
		AllowHeaders:     []string{"Origin", "Content-Type"},       // 許可するヘッダー
		AllowCredentials: true,                                     // Cookieを許可
		MaxAge:           12 * 3600,                                // CORSのキャッシュ時間
	}))

	router.SetupRoutes(r)
	r.Run() // デフォルトは :8080
}
