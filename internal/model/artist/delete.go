package artist_model

import "gorm.io/gorm"

// DeleteArtist deletes an existing artist from the database
func DeleteArtist(db *gorm.DB, artist_id int) error {
	return db.Delete(&Artist{}, artist_id).Error
}

// DeleteArtists deletes existing artists from the database
func DeleteArtists(db *gorm.DB, artist_ids []int) error {
	return db.Where("Id IN ?", artist_ids).Delete(&Artist{}).Error
}
