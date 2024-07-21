package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

// PostPlaylist creates a new playlist in the database
func CreatePlaylist(db *gorm.DB, title string, owners_ptr []*db_tables.User, description string, illustration string) (int, error) {
	owners := make([]db_tables.User, len(owners_ptr))
	for i, owner := range owners_ptr {
		owners[i] = *owner
	}
	new_playlist := db_tables.Playlist{
		Title:        title,
		Illustration: illustration,
		Description:  description,
		Owners:       owners,
		Subscribers:  owners,
	}
	err := db.Create(&new_playlist).Error
	if err != nil {
		return -1, err
	}
	return new_playlist.Id, nil
}
