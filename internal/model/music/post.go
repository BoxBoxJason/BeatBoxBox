package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// CreateMusic creates a new music in the database
func CreateMusic(db *gorm.DB, title string, artist_id int, genres []string, album_id int, file_name string, illustration_path string) error {
	new_music := db_model.Music{
		Title:        title,
		Path:         file_name,
		AlbumId:      album_id,
		Illustration: illustration_path,
	}
	return db.Create(&new_music).Error
}
