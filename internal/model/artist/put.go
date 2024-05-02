package artist_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// UpdateMusic updates an existing music in the database
func UpdateArtist(db *gorm.DB, artist_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Artist{}).Where("Id = ?", artist_id).Updates(update_map).Error
}
