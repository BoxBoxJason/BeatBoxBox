package db_controller

import (
	"BeatBoxBox/pkg/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Musics struct {
	id       int     `gorm:"primaryKey;autoIncrement"`
	titre    string  `gorm:"type:text;not null"`
	auteur   string  `gorm:"type:text;not null"`
	album    string  `gorm:"type:text"`
	genre    string  `gorm:"type:text;not null"`
	nbecoute int     `gorm:"type:int;not null"`
	note     float32 `gorm:"type:float32;not null"`
	nbnotes  int     `gorm:"type:int;not null"`
	path     string  `gorm:"type:varchar(255);not null"`
}

type Users struct {
	pseudo                  string `gorm:"type:varchar(32);not null"`
	subscribed_playlist_ids string `gorm:"type:varchar;not null"`
	liked_music_ids         string `gorm:"type:varchar;not null"`
	hashed_password         string `gorm:"type:varchar(64);not null"`
	salt                    string `gorm:"type:varchar(16);not null"`
	id                      int    `gorm:"primaryKey;autoIncrement"`
}

type Playlists struct {
	titre     string `gorm:"type:text;not null"`
	music_ids string `gorm:"type:text"`
	id        int    `gorm:"primaryKey;autoIncrement"`
	createur  int    `gorm:"foreignKey;autoIncrement"`
}

func CreateDB() {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database")
	}

	// Auto-migrate the schema
	db.AutoMigrate(&Users{}, &Musics{}, &Playlists{})

	logger.Info("Tables created successfully")
}