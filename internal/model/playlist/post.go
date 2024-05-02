package playlist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostPlaylist creates a new playlist in the database
func CreatePlaylist(db *gorm.DB, title string, creator_id int, illustration string) error {
	new_playlist := db_model.Playlist{
		Title:        title,
		CreatorId:    creator_id,
		Illustration: illustration,
	}
	return db.Create(&new_playlist).Error
}
