package music_model

import "gorm.io/gorm"

// DELETE METHODS

// DeleteMusic deletes an existing music from the database
func DeleteMusic(db *gorm.DB, music_id int) error {
	// Deletes the music entry where the Id matches the provided music_id
	result := db.Delete(&Music{}, music_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteMusics deletes existing musics from the database
func DeleteMusics(db *gorm.DB, music_ids []int) error {
	// Deletes music entries where their Id is in the provided music_ids slice
	result := db.Where("Id IN ?", music_ids).Delete(&Music{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
