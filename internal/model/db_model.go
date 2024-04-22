package db_model

import (
	"BeatBoxBox/pkg/logger"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open a database connection to the PostgreSQL database
func OpenDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Critical("failed to connect database")
		return nil, err
	}
	return db, nil
}

// Close a database connection to the PostgreSQL database
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Critical("failed to get database connection")
		return
	}
	sqlDB.Close()
}
