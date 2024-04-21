package db_controller

import (
	"BeatBoxBox/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Pseudo               string     `gorm:"type:varchar(32);unique;not null"`
	Subscribed_playlists []Playlist `gorm:"foreignKey:PlaylistIds"`
	PlaylistIds          []int
	Liked_musics         []Music `gorm:"foreignKey:MusicIds"`
	MusicIds             []int
	Hashed_password      string `gorm:"type:varchar(64);not null"`
	Salt                 string `gorm:"type:varchar(16);not null"`
	Id                   int    `gorm:"primaryKey;autoIncrement"`
}

type Artist struct {
	Pseudo string `gorm:"type:varchar(32);unique;not null"`
	Id     int    `gorm:"primaryKey;autoIncrement"`
}

type Music struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Title      string `gorm:"type:text;not null"`
	ArtistId   int
	Artist     Artist `gorm:"foreignKey:ArtistId"`
	AlbumId    int
	Album      Album   `gorm:"foreignKey:AlbumId"`
	Genres     string  `gorm:"type:text"`
	Nblistened int     `gorm:"type:int;default 0"`
	Rating     float32 `gorm:"type:float32;default 0"`
	Nbrating   int     `gorm:"type:int;default 0"`
	Path       string  `gorm:"type:varchar(255);not null"`
}

type Playlist struct {
	Title     string `gorm:"type:text;unique;not null"`
	Musics    []Music
	Id        int `gorm:"primaryKey;autoIncrement"`
	CreatorId int
	Creator   User `gorm:"foreignKey:UserId"`
}

type Album struct {
	Title    string `gorm:"type:text;unique;not null"`
	MusicIds []int
	Musics   []Music `gorm:"foreignKey:MusicIds"`
	Id       int     `gorm:"primaryKey;autoIncrement"`
	ArtistId int
	Artist   Artist `gorm:"foreignKey:ArtistId"`
}

func CreateDB() {
	// Connect to the database
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Music{}, &Playlist{}, &Album{}, &Artist{})
	logger.Info("Tables created successfully")

}
