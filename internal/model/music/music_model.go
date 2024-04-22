package music_model

import (
	artist_model "BeatBoxBox/internal/model/artist"
	user_model "BeatBoxBox/internal/model/user"
)

type Music struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Title      string `gorm:"type:text;not null"`
	ArtistId   int
	Artist     artist_model.Artist `gorm:"foreignKey:ArtistId"`
	Genres     string              `gorm:"type:text"`
	Nblistened int                 `gorm:"default:0"`
	Rating     float32             `gorm:"default:0"`
	Nbrating   int                 `gorm:"default:0"`
	Path       string              `gorm:"type:varchar(255);not null"`
	UploaderId int
	Uploader   user_model.User `gorm:"foreignKey:UploaderId"`
}
