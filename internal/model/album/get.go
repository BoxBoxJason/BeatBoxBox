package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// GetAlbumsFromFilters returns a list of albums from the database
// Filters can be passed to filter the albums
func GetAlbumsFromFilters(db *gorm.DB, filters map[string]interface{}) []db_tables.Album {
	raw_albums := db_model.GetRecordsByFields(db, &db_tables.Album{}, filters)
	if raw_albums == nil {
		return nil
	}
	albums := make([]db_tables.Album, len(raw_albums))
	for i, album := range raw_albums {
		albums[i] = album.(db_tables.Album)
	}

	return albums
}

// GetAlbum returns a album from the database
// Selects the album with the given album_id
func GetAlbum(db *gorm.DB, album_id int) (db_tables.Album, error) {
	album := db_model.GetRecordFromId(db, &db_tables.Album{}, album_id)
	if album == nil {
		return db_tables.Album{}, gorm.ErrRecordNotFound
	}
	return *album.(*db_tables.Album), nil
}

// GetAlbums returns a list of albums from the database
// Selects the albums with the given album_ids
func GetAlbums(db *gorm.DB, album_ids []int) ([]db_tables.Album, error) {
	records := db_model.GetRecordsFromIds(db, &db_tables.Album{}, album_ids)
	if records == nil {
		return nil, gorm.ErrRecordNotFound
	}
	albums := make([]db_tables.Album, len(records))
	for i, record := range records {
		albums[i] = record.(db_tables.Album)
	}

	return albums, nil
}

// GetAlbumsFromPartialTitle returns a list of albums from the database
// Selects the albums with the given partial_title
func GetAlbumsFromPartialTitle(db *gorm.DB, filters map[string]interface{}, partial_title string) []db_tables.Album {
	records := db_model.GetRecordsByFieldsWithCondition(db, &db_tables.Album{}, filters, "title LIKE ?", "%"+partial_title+"%")
	if records == nil {
		return nil
	}
	albums := make([]db_tables.Album, len(records))
	for i, record := range records {
		albums[i] = record.(db_tables.Album)
	}

	return albums
}
