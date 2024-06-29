package music_model

import (
	db_tables "BeatBoxBox/internal/model"
	db_model "BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// UpdateMusic updates an existing music in the database
func UpdateMusic(db *gorm.DB, music *db_tables.Music, fields map[string]interface{}) error {
	return db.Model(music).Updates(fields).Error
}

func AddArtistsToMusic(db *gorm.DB, music *db_tables.Music, artists []*db_tables.Artist) error {
	return db_model.AddElementsToAssociation(db, music, "Artists", artists)
}

func RemoveArtistsFromMusic(db *gorm.DB, music *db_tables.Music, artists []*db_tables.Artist) error {
	return db_model.RemoveElementsFromAssociation(db, music, "Artists", artists)
}
