package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

// CreateArtist creates a new artist in the database
func CreateArtist(db *gorm.DB, pseudo string, genres []string, bio string, birthdate string, illustration string) (db_tables.Artist, error) {
	new_artist := db_tables.Artist{
		Pseudo:       pseudo,
		Bio:          bio,
		Illustration: illustration,
		BirthDate:    birthdate,
		Genres:       genres,
	}
	err := db.Create(&new_artist).Error
	if err != nil {
		return db_tables.Artist{}, err
	}
	return new_artist, nil
}
