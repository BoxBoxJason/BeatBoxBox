package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// DeleteAlbumFromRecord is a function that deletes an album from the database.
func DeleteAlbumFromRecord(db *gorm.DB, album db_tables.Album) error {
	return db_model.DeleteDBRecordNoFetch(db, &album)
}

// DeleteAlbumsFromRecords is a function that deletes multiple albums from the database.
func DeleteAlbumsFromRecords(db *gorm.DB, albums []*db_tables.Album) error {
	return db.Delete(albums).Error
}
