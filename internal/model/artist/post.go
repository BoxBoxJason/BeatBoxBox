package artist_model

import "gorm.io/gorm"

// CreateArtist creates a new artist in the database
func CreateArtist(db *gorm.DB, pseudo string, illustration string) error {
	new_artist := Artist{
		Pseudo:       pseudo,
		Illustration: illustration,
	}
	return db.Create(&new_artist).Error
}
