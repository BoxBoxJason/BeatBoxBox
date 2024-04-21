package album_model

import (
	artist_model "BeatBoxBox/internal/model/artist"
	music_model "BeatBoxBox/internal/model/music"
)

type Album struct {
	Title    string `gorm:"type:text;not null"`
	MusicIds []int
	Musics   []music_model.Music `gorm:"foreignKey:MusicIds"`
	Id       int                 `gorm:"primaryKey;autoIncrement"`
	ArtistId int
	Artist   artist_model.Artist `gorm:"foreignKey:ArtistId"`
}
