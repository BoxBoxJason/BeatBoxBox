package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

func UpdateAlbum(db *gorm.DB, album *db_tables.Album, update_map map[string]interface{}) error {
	return db_model.EditRecordFields(db, album, update_map)
}

func AddMusicsToAlbum(db *gorm.DB, album *db_tables.Album, musics []*db_tables.Music) error {
	return db_model.AddElementsToAssociation(db, album, "Musics", musics)
}

func RemoveMusicsFromAlbum(db *gorm.DB, album *db_tables.Album, musics []*db_tables.Music) error {
	return db_model.RemoveElementsFromAssociation(db, album, "Musics", musics)
}

func AddArtistsToAlbum(db *gorm.DB, album *db_tables.Album, artists []*db_tables.Artist) error {
	return db_model.AddElementsToAssociation(db, album, "Artists", artists)
}

func RemoveArtistsFromAlbum(db *gorm.DB, album *db_tables.Album, artists []*db_tables.Artist) error {
	return db_model.RemoveElementsFromAssociation(db, album, "Artists", artists)
}
