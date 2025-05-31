package main

import (
    "github.com/gin-gonic/gin"
    "paper-app-backend/internal/router"
    "paper-app-backend/internal/db"
)

func main() {
    // データベースの初期化
    db.InitDB()

    // ルーターのセットアップ
    r := gin.Default()
    router.SetupRoutes(r)
    r.Run() // デフォルトは :8080
}