package playlist_model

import "gorm.io/gorm"

// UpdatePlaylist updates an existing playlist in the database
func UpdatePlaylist(db *gorm.DB, playlist_id int, update_map map[string]interface{}) error {
	return db.Model(&Playlist{}).Where("Id = ?", playlist_id).Updates(update_map).Error
}
