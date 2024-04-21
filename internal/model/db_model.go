package db_model

import (
	album_model "BeatBoxBox/internal/model/album"
	artist_model "BeatBoxBox/internal/model/artist"
	music_model "BeatBoxBox/internal/model/music"
	playlist_model "BeatBoxBox/internal/model/playlist"
	user_model "BeatBoxBox/internal/model/user"
	"BeatBoxBox/pkg/logger"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	db, err := openDB()
	if err != nil {
		logger.Critical("Failed to connect database: ", err)
	} else {
		db.AutoMigrate(&user_model.User{}, &music_model.Music{}, &playlist_model.Playlist{}, &album_model.Album{}, &artist_model.Artist{})
		logger.Info("Tables created successfully")
	}
}

// Open a database connection to the PostgreSQL database
func openDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Critical("failed to connect database")
		return nil, err
	}
	return db, nil
}