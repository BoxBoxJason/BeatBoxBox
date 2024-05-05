package album_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

func UpdateAlbum(db *gorm.DB, album_id int, update_map map[string]interface{}) error {
	return db.Model(&db_model.Album{}).Where("id = ?", album_id).Updates(update_map).Error
}
