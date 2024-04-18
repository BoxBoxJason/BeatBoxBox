package db_controller

import (
	"BeatBoxBox/pkg/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Musics struct {
	id         int     `gorm:"primaryKey;autoIncrement"`
	title      string  `gorm:"type:text;unique;not null"`
	author     string  `gorm:"type:text;not null"`
	album      string  `gorm:"type:text"`
	genre      string  `gorm:"type:text"`
	nblistened int     `gorm:"type:int;default 0"`
	rating     float32 `gorm:"type:float32;default 0"`
	nbrating   int     `gorm:"type:int;default 0"`
	path       string  `gorm:"type:varchar(255);not null"`
}

type Users struct {
	pseudo                  string `gorm:"type:varchar(32);unique;not null"`
	subscribed_playlist_ids string `gorm:"type:varchar;not null"`
	liked_music_ids         string `gorm:"type:varchar;not null"`
	hashed_password         string `gorm:"type:varchar(64);not null"`
	salt                    string `gorm:"type:varchar(16);not null"`
	id                      int    `gorm:"primaryKey;autoIncrement"`
}

type Playlists struct {
	title     string `gorm:"type:text;unique;not null"`
	music_ids string `gorm:"type:text"`
	id        int    `gorm:"primaryKey;autoIncrement"`
	creator   int    `gorm:"foreignKey"`
}

func CreateDB() {
	// Connect to the database
	_, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database")
	}

	logger.Info("Tables created successfully")
}
