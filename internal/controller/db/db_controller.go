package db_controller

import (
	"BeatBoxBox/pkg/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Musics struct {
	Id         int     `gorm:"primaryKey;autoIncrement"`
	Title      string  `gorm:"type:text;unique;not null"`
	Author     string  `gorm:"type:text;not null"`
	Album      string  `gorm:"type:text"`
	Genre      string  `gorm:"type:text"`
	Nblistened int     `gorm:"type:int;default 0"`
	Rating     float32 `gorm:"type:float32;default 0"`
	Nbrating   int     `gorm:"type:int;default 0"`
	Path       string  `gorm:"type:varchar(255);not null"`
}

type Playlists struct {
	Title     string `gorm:"type:text;unique;not null"`
	Music_ids string `gorm:"type:text"`
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Creator   int    `gorm:"foreignKey"`
}

func CreateDB() {
	type Users struct {
		Pseudo                  string `gorm:"type:varchar(32);unique;not null"`
		Subscribed_playlist_ids string `gorm:"type:varchar;not null"`
		Liked_music_ids         string `gorm:"type:varchar;not null"`
		Hashed_password         string `gorm:"type:varchar(64);not null"`
		Salt                    string `gorm:"type:varchar(16);not null"`
		Id                      int    `gorm:"primaryKey;autoIncrement"`
	}
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database")
	}

	db.AutoMigrate(&Users{})

	user := Users{Pseudo: "bob", Subscribed_playlist_ids: "", Liked_music_ids: "", Hashed_password: "1234", Salt: "1234", Id: 1}

	db.Create(&user)

	logger.Info("Tables created successfully")
}
