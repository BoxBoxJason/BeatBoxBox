package playlist_model

import "gorm.io/gorm"

// PostPlaylist creates a new playlist in the database
func CreatePlaylist(db *gorm.DB, title string, creator_id int, illustration string) error {
	new_playlist := Playlist{
		Title:        title,
		CreatorId:    creator_id,
		Illustration: illustration,
	}
	return db.Create(&new_playlist).Error
}
