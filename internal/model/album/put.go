package album_model

import "gorm.io/gorm"

func UpdateMusic(db *gorm.DB, music_id int, update_map map[string]interface{}) error {
	return db.Model(&Album{}).Where("Id = ?", music_id).Updates(update_map).Error
}
