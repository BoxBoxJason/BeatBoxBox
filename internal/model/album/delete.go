package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// DeleteAlbum is a function that deletes an album from the database.
func DeleteAlbum(db *gorm.DB, album_id int) error {
	return db_model.DeleteDBRecord(db, &db_tables.Album{}, album_id)
}

// DeleteAlbums is a function that deletes multiple albums from the database.
func DeleteAlbums(db *gorm.DB, album_ids []int) error {
	return db_model.DeleteDBRecords(db, &db_tables.Album{}, album_ids)
}
