package playlist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// UpdatePlaylist updates an existing playlist in the database
func UpdatePlaylist(db *gorm.DB, playlist_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Playlist{}).Where("id = ?", playlist_id).Updates(update_map).Error
}
