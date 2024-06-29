package db_model

import (
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/logger"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open a database connection to the PostgreSQL database
func OpenDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	if host == "" || port == "" || user == "" || password == "" || dbname == "" || sslmode == "" {
		logger.Critical("Missing environment variables")
		return nil, custom_errors.NewDatabaseError("Missing at least one of the required environment variables: DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Critical("failed to connect database")
		return nil, custom_errors.NewDatabaseError(err.Error())
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

// Checks if the database connection is alive
func CheckDB() error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	return nil
}
