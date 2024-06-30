package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"gorm.io/gorm"
)

// DeleteArtistFromRecord deletes an existing artist from the database
func DeleteArtistFromRecord(db *gorm.DB, artist db_tables.Artist) error {
	return db.Delete(&artist).Error
}

// DeleteArtists deletes existing artists from the database
func DeleteArtistsFromRecords(db *gorm.DB, artists []*db_tables.Artist) error {
	return db.Delete(artists).Error
}
