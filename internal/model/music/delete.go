package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// DeleteMusic deletes an existing music from the database
func DeleteMusic(db *gorm.DB, music_id int) error {
	return db.Delete(&db_model.Music{}, music_id).Error
}

// DeleteMusics deletes existing musics from the database
func DeleteMusics(db *gorm.DB, music_ids []int) error {
	return db.Where("id IN ?", music_ids).Delete(&db_model.Music{}).Error
}
