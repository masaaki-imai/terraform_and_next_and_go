package database

import (
	"L-cart/models"
	"fmt"
	"log"
	"os"

	"L-cart/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo"
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// 作成したデータベースに接続
	finalDsn := fmt.Sprintf(dsn, dbUser, dbPass, dbHost, dbName)
	db, err = gorm.Open(mysql.Open(finalDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	return db
}

func AutoMigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		logger.LogError(err, "Failed to migrate")
		log.Fatalf("Failed to migrate: %v", err)
	} else {
		logger.LogInfo("Database migration completed successfully")
		log.Println("Database migration completed successfully")
	}
}
