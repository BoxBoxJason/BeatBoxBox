package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// UpdateMusic updates an existing music in the database
func UpdateMusic(db *gorm.DB, music_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Music{}).Where("Id = ?", music_id).Updates(update_map).Error
}
