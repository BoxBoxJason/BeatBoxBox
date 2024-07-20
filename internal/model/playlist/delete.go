package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// DeletePlaylist is a function that deletes an existing playlist from the database.
func DeletePlaylist(db *gorm.DB, playlist *db_tables.Playlist) error {
	return db_model.DeleteDBRecordNoFetch(db, playlist)
}

// DeletePlaylists is a function that deletes multiple existing playlists from the database.
func DeletePlaylists(db *gorm.DB, playlists []*db_tables.Playlist) error {
	return db.Delete(playlists).Error
}
