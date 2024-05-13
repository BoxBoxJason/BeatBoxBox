package playlist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// PostPlaylist creates a new playlist in the database
// Returns corresponding Id
func CreatePlaylist(db *gorm.DB, title string, creator_id int, description string, illustration string) (int, error) {
	new_playlist := db_model.Playlist{
		Title:        title,
		CreatorId:    creator_id,
		Illustration: illustration,
		Description:  description,
	}
	err := db.Create(&new_playlist).Error
	if err != nil {
		return -1, err
	}
	return new_playlist.Id, nil
}
