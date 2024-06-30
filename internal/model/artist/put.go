package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// UpdateMusic updates an existing music in the database
func UpdateArtist(db *gorm.DB, artist *db_tables.Artist, update_map map[string]interface{}) error {
	return db_model.EditRecordFields(db, artist, update_map)
}

// AddMusicsToArtist adds musics to an artist in the database
func AddMusicsToArtist(db *gorm.DB, artist *db_tables.Artist, musics []*db_tables.Music) error {
	return db_model.AddElementsToAssociation(db, artist, "Musics", musics)
}

// RemoveMusicsFromArtist removes musics from an artist in the database
func RemoveMusicsFromArtist(db *gorm.DB, artist *db_tables.Artist, musics []*db_tables.Music) error {
	return db_model.RemoveElementsFromAssociation(db, artist, "Musics", musics)
}
