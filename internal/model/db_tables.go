package db_tables

import (
	"BeatBoxBox/pkg/db_model"
	"BeatBoxBox/pkg/logger"
)

// Initialize the database connection and create the tables
func init() {
	db, err := db_model.OpenDB()
	if err != nil {
		logger.Critical("Failed to connect database: ", err)
	} else {
		db.AutoMigrate(Artist{}, User{}, Album{}, Music{}, Playlist{}, AuthCookie{})
		logger.Info("Tables created successfully")
	}
}

type Playlist struct {
	Id           int     `gorm:"primaryKey;autoIncrement"`
	Title        string  `gorm:"type:text;unique;not null"`
	Description  string  `gorm:"type:text"`
	Illustration string  `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music `gorm:"many2many:playlist_musics;"`
	CreatorId    int
	Creator      User `gorm:"foreignKey:CreatorId"`
	Protected    bool `gorm:"default:true"`
}

type Album struct {
	Id           int      `gorm:"primaryKey;autoIncrement"`
	Title        string   `gorm:"type:text;not null"`
	Description  string   `gorm:"type:text"`
	Artists      []Artist `gorm:"many2many:album_artists;"`
	Illustration string   `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music  `gorm:"many2many:album_musics;"`
}

type Artist struct {
	Id           int     `gorm:"primaryKey;autoIncrement"`
	Pseudo       string  `gorm:"type:varchar(32);unique;not null"`
	Bio          string  `gorm:"type:text"`
	Illustration string  `gorm:"type:text;default:'default.jpg'"`
	Musics       []Music `gorm:"foreignKey:AlbumId;"`
}

type Music struct {
	Id           int      `gorm:"primaryKey;autoIncrement"`
	Title        string   `gorm:"type:text;not null"`
	Lyrics       string   `gorm:"type:text"`
	Artists      []Artist `gorm:"many2many:artist_musics;"`
	AlbumId      int
	Album        Album    `gorm:"foreignKey:AlbumId"`
	Genres       []string `gorm:"type:text"`
	Nblistened   int      `gorm:"default:0"`
	Rating       float32  `gorm:"default:0"`
	Nbrating     int      `gorm:"default:0"`
	Likes        int      `gorm:"default:0"`
	Path         string   `gorm:"type:varchar(36);not null"`
	Illustration string   `gorm:"type:text;default:'default.jpg'"`
	UploaderId   int
	Uploader     User `gorm:"foreignKey:UploaderId"`
}

type User struct {
	Pseudo              string     `gorm:"type:varchar(32);unique;not null"`
	Email               string     `gorm:"type:varchar(256);unique;not null"`
	Hashed_password     string     `gorm:"type:varchar(64);not null"`
	Id                  int        `gorm:"primaryKey;autoIncrement"`
	Illustration        string     `gorm:"type:text;default:'default.jpg'"`
	SubscribedPlaylists []Playlist `gorm:"many2many:user_subscribed_playlists;"`
	Playlists           []Playlist `gorm:"foreignKey:CreatorId"`
	LikedMusics         []Music    `gorm:"many2many:user_liked_musics;"`
	UploadedMusics      []Music    `gorm:"foreignKey:UploaderId"`
}

type AuthCookie struct {
	Id              int    `gorm:"primaryKey;autoIncrement"`
	HashedAuthToken string `gorm:"type:varchar(60);not null"`
	UserId          int
	User            User  `gorm:"foreignKey:UserId"`
	ExpirationDate  int64 `gorm:"not null"`
}
