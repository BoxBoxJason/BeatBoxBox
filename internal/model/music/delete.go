package music_model

import "gorm.io/gorm"

// DeleteMusic deletes an existing music from the database
func DeleteMusic(db *gorm.DB, music_id int) error {
	return db.Delete(&Music{}, music_id).Error
}

// DeleteMusics deletes existing musics from the database
func DeleteMusics(db *gorm.DB, music_ids []int) error {
	return db.Where("Id IN ?", music_ids).Delete(&Music{}).Error
}
