package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

// PostPlaylist creates a new playlist in the database
// Returns corresponding Id
func CreatePlaylist(db *gorm.DB, title string, creator_id int, description string, illustration string) (int, error) {
	new_playlist := db_tables.Playlist{
		Title:        title,
		Illustration: illustration,
		Description:  description,
	}
	if creator_id >= 0 {
		creator_id_uint := uint(creator_id)
		new_playlist.CreatorId = &creator_id_uint
	}
	err := db.Create(&new_playlist).Error
	if err != nil {
		return -1, err
	}
	return new_playlist.Id, nil
}
