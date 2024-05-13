package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// CreateMusic creates a new music in the database
func CreateMusic(db *gorm.DB, title string, artist_id int, genres []string, album_id int, file_name string, illustration_path string) (int, error) {
	new_music := db_model.Music{
		Title:        title,
		Path:         file_name,
		AlbumId:      album_id,
		Illustration: illustration_path,
	}
	err := db.Create(&new_music).Error
	if err != nil {
		return -1, err
	}
	return new_music.Id, nil
}
