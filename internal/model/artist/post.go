package artist_model

import "gorm.io/gorm"

// CreateArtist creates a new artist in the database
func CreateArtist(db *gorm.DB, pseudo string) error {
	new_artist := Artist{
		Pseudo: pseudo,
	}
	return db.Create(&new_artist).Error
}
