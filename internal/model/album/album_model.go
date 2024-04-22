package album_model

import (
	artist_model "BeatBoxBox/internal/model/artist"
)

type Album struct {
	Title    string `gorm:"type:text;not null"`
	Id       int    `gorm:"primaryKey;autoIncrement"`
	ArtistId int
	Artist   artist_model.Artist `gorm:"foreignKey:ArtistId"`
}
