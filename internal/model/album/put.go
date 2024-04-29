package album_model

import "gorm.io/gorm"

func UpdateAlbum(db *gorm.DB, album_id int, update_map map[string]interface{}) error {
	return db.Model(&Album{}).Where("Id = ?", album_id).Updates(update_map).Error
}
