package music_model

import (
	"strings"

	"gorm.io/gorm"
)

// POST METHODS

// CreateMusic creates a new music in the database
func CreateMusic(db *gorm.DB, title string, artist_id int, genres []string, album_id int, file_name string) error {
	new_music := Music{
		Title:    title,
		ArtistId: artist_id,
		Genres:   strings.Join(genres, ","),
		Path:     file_name,
		AlbumId:  album_id,
	}
	result := db.Create(&new_music)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
