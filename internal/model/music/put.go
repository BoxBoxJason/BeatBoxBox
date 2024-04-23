package music_model

import "gorm.io/gorm"

// UpdateMusic updates an existing music in the database
func UpdateMusic(db *gorm.DB, music_id int, update_map map[string]interface{}) error {
	return db.Model(&Music{}).Where("Id = ?", music_id).Updates(update_map).Error
}
