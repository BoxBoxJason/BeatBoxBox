package playlist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// DeletePlaylist is a function that deletes an playlist from the database.
func DeletePlaylist(db *gorm.DB, playlist_id int) error {
	return db.Delete(&db_model.Playlist{}, playlist_id).Error
}

// DeletePlaylists is a function that deletes multiple playlists from the database.
func DeletePlaylists(db *gorm.DB, playlist_ids []int) error {
	return db.Where("Id IN ?", playlist_ids).Delete(&db_model.Playlist{}).Error
}
