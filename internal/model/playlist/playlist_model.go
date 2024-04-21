package playlist_model

import (
	music_model "BeatBoxBox/internal/model/music"
	user_model "BeatBoxBox/internal/model/user"
)

type Playlist struct {
	Title     string              `gorm:"type:text;unique;not null"`
	Musics    []music_model.Music `gorm:"foreignKey:MusicIds"`
	Id        int                 `gorm:"primaryKey;autoIncrement"`
	CreatorId int
	Creator   user_model.User `gorm:"foreignKey:UserId"`
}
