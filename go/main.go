package main

import (
	"L-cart/database"
	"L-cart/routes"
	"L-cart/translations"
	"L-cart/utils/logger"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// ロガーの初期化
	logger.Init()
	database.ConnectDB()
	database.AutoMigrateDB()

	// 日本時間のロケーションを設定（AWSの環境では名前での設定が失敗するためオフセットを使用）
	jst := time.FixedZone("JST", 9*60*60)

	// デフォルトのタイムゾーンとして設定
	time.Local = jst

	// バリデーターの初期化
	translations.InitValidator()

	// Seed test data for development
	if _, err := database.SeedUser(); err != nil {
		log.Printf("Warning: Failed to seed user: %v", err)
	}

	// Ginのログフォーマットをカスタマイズ
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.LogInfo(fmt.Sprintf(
			"[GIN-debug] %-6s %-40s --> %s (%d handlers)",
			httpMethod,
			absolutePath,
			handlerName,
			nuHandlers,
		))
	}

	// カスタムログフォーマッターを設定
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return logger.FormatGinLog(param)
	}))
	r.Use(gin.Recovery())

	routes.SetupRoutes(r)

	r.Run(":8080")
}
