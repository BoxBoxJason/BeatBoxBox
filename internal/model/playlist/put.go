package playlist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// UpdatePlaylist updates an existing playlist in the database
func UpdatePlaylist(db *gorm.DB, playlist *db_tables.Playlist, update_map map[string]interface{}) error {
	return db_model.EditRecordFields(db, playlist, update_map)
}

func AddMusicsToPlaylist(db *gorm.DB, playlist *db_tables.Playlist, musics []*db_tables.Music) error {
	return db_model.AddElementsToAssociation(db, playlist, "Musics", musics)
}

func RemoveMusicsFromPlaylist(db *gorm.DB, playlist *db_tables.Playlist, musics []*db_tables.Music) error {
	return db_model.RemoveElementsFromAssociation(db, playlist, "Musics", musics)
}
