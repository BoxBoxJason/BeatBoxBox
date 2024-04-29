package playlist_model

import (
	user_model "BeatBoxBox/internal/model/user"
)

type Playlist struct {
	Title        string `gorm:"type:text;unique;not null"`
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Illustration string `gorm:"type:varchar(36);default:'default.jpg'"`
	CreatorId    int
	Creator      user_model.User `gorm:"foreignKey:CreatorId"`
}
