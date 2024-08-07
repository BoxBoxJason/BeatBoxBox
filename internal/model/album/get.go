package album_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// GetAlbumsFromFilters returns a list of albums from the database
// Filters can be passed to filter the albums
func GetAlbumsFromFilters(db *gorm.DB, titles []string, partial_titles []string, genres []string, artists_names []string, musics_names []string, artists_ids []int, musics_ids []int) []db_tables.Album {
	var albums []db_tables.Album
	query := db.Model(&db_tables.Album{})
	if len(genres) > 0 {
		for _, genre := range genres {
			query = query.Where("genres @> ARRAY[?]::text[]", genre)
		}
	}
	if len(titles) > 0 {
		query = query.Where("title IN ?", titles)
	} else if len(partial_titles) > 0 {
		for _, partial_title := range partial_titles {
			query = query.Or("title LIKE ?", "%"+partial_title+"%")
		}
	}
	if len(artists_names) > 0 {
		query = query.Joins("JOIN album_artists ON album_artists.album_id = albums.id").
			Joins("JOIN artists ON artists.id = album_artists.artist_id")
		for _, artist_name := range artists_names {
			query = query.Where("artists.pseudo = ?", artist_name)
		}
	} else if len(artists_ids) > 0 {
		db.Joins("JOIN album_artists ON album_artists.album_id = albums.id")
		for _, artist_id := range artists_ids {
			db.Where("album_artists.artist_id ", artist_id)
		}
	}
	if len(musics_names) > 0 {
		db.Joins("JOIN musics ON musics.album_id = albums.id")
		for _, music_name := range musics_names {
			db.Where("musics.title = ?", music_name)
		}
	} else if len(musics_ids) > 0 {
		db.Joins("JOIN musics ON musics.album_id = albums.id")
		for _, music_id := range musics_ids {
			db.Where("musics.id = ?", music_id)
		}
	}
	db.Group("albums.id").Find(&albums)
	return albums
}

// GetAlbum returns an album from the database
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

func AlbumAlreadyExists(db *gorm.DB, title string, artists_ids []int) bool {
	if len(artists_ids) == 0 {
		return false
	}
	var album db_tables.Album
	err := db.Where("title = ?", title).
		Joins("JOIN album_artists ON album_artists.album_id = albums.id").
		Where("album_artists.artist_id IN ?", artists_ids).
		Group("albums.id").
		Having("COUNT(DISTINCT album_artists.artist_id) = ?", len(artists_ids)).
		First(&album).Error
	return err == nil
}
