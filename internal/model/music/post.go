package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// CreateMusic creates a new music in the database
func CreateMusic(db *gorm.DB, title string, genres []string, album_id int, file_name string, illustration_path string, uploader_id int) (int, error) {
	new_music := db_model.Music{}
	new_music = db_model.Music{
		Title:        title,
		Path:         file_name,
		Genres:       genres,
		Illustration: illustration_path,
	}
	if album_id >= 0 {
		album_id_uint := uint(album_id)
		new_music.AlbumId = &album_id_uint
	}
	if uploader_id >= 0 {
		uploader_id_uint := uint(uploader_id)
		new_music.UploaderId = &uploader_id_uint
	}
	err := db.Create(&new_music).Error
	if err != nil {
		return -1, err
	}
	return new_music.Id, nil
}
