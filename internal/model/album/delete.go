package album_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// DeleteAlbum is a function that deletes an album from the database.
func DeleteAlbum(db *gorm.DB, album_id int) error {
	return db.Delete(&db_model.Album{}, album_id).Error
}

// DeleteAlbums is a function that deletes multiple albums from the database.
func DeleteAlbums(db *gorm.DB, album_ids []int) error {
	return db.Where("Id IN ?", album_ids).Delete(&db_model.Album{}).Error
}
