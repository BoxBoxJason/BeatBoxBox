package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

// CreateArtist creates a new artist in the database
func CreateArtist(db *gorm.DB, pseudo string, bio string, illustration string) (int, error) {
	new_artist := db_tables.Artist{
		Pseudo:       pseudo,
		Bio:          bio,
		Illustration: illustration,
	}
	err := db.Create(&new_artist).Error
	if err != nil {
		return -1, err
	}
	return new_artist.Id, nil
}
