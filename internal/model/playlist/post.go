package playlist_model

import "gorm.io/gorm"

// PostPlaylist creates a new playlist in the database
func PostPlaylist(db *gorm.DB, title string, creator_id int) error {
	new_playlist := Playlist{
		Title:     title,
		CreatorId: creator_id,
	}
	return db.Create(&new_playlist).Error
}
