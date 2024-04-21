/*
package music_model is the model for the musics.

Handles the database connection and the CRUD operations for the musics.
*/

package music_model

import (
	artist_model "BeatBoxBox/internal/model/artist"
)

type Music struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	Title      string `gorm:"type:text;not null"`
	ArtistId   int
	Artist     artist_model.Artist `gorm:"foreignKey:ArtistId"`
	Genres     string              `gorm:"type:text"`
	Nblistened int                 `gorm:"type:int;default 0"`
	Rating     float32             `gorm:"type:float32;default 0"`
	Nbrating   int                 `gorm:"type:int;default 0"`
	Path       string              `gorm:"type:varchar(255);not null"`
}
