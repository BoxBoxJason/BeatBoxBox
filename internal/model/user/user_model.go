package user_model

import (
	music_model "BeatBoxBox/internal/model/music"
)

type User struct {
	Pseudo                string `gorm:"type:varchar(32);unique;not null"`
	SubscribedPlaylistIds []int
	Liked_musics          []music_model.Music `gorm:"foreignKey:MusicIds"`
	MusicIds              []int
	Hashed_password       string `gorm:"type:varchar(64);not null"`
	Salt                  string `gorm:"type:varchar(16);not null"`
	Id                    int    `gorm:"primaryKey;autoIncrement"`
}
