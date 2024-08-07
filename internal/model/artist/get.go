package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"

	"gorm.io/gorm"
)

// GetArtistsFromFilters returns a list of artists from the database
// Filters can be passed to filter the artists
func GetArtistsFromFilters(db *gorm.DB, pseudos []string, partial_pseudos []string, genres []string, albums_ids []int, albums []string, musics_ids []int, musics []string) []db_tables.Artist {
	var artists []db_tables.Artist
	query := db.Model(&db_tables.Artist{})
	if len(genres) > 0 {
		for _, genre := range genres {
			query = query.Where("genres @> ARRAY[?]::text[]", genre)
		}
	}
	if len(pseudos) > 0 {
		query = query.Where("pseudo IN ?", pseudos)
	} else if len(partial_pseudos) > 0 {
		for _, partial_pseudo := range partial_pseudos {
			query = query.Or("pseudo LIKE ?", "%"+partial_pseudo+"%")
		}
	}
	if len(albums_ids) > 0 {
		query = query.Joins("JOIN album_artists ON album_artists.artist_id = artists.id").
			Where("album_artists.album_id IN ?", albums_ids)
	} else if len(albums) > 0 {
		query = query.Joins("JOIN album_artists ON album_artists.artist_id = artists.id").
			Joins("JOIN albums ON albums.id = album_artists.album_id").
			Where("albums.title IN ?", albums)
	}
	if len(musics_ids) > 0 {
		query = query.Joins("JOIN musics ON musics.album_id = albums.id").
			Where("musics.id IN ?", musics_ids)
	} else if len(musics) > 0 {
		query = query.Joins("JOIN musics ON musics.album_id = albums.id").
			Where("musics.title IN ?", musics)
	}
	query.Group("artists.id").Find(&artists)
	return artists
}

// GetArtist returns a artist from the database
// Selects the artist with the given artist_id
func GetArtist(db *gorm.DB, artist_id int) (db_tables.Artist, error) {
	artist := db_model.GetRecordFromId(db, &db_tables.Artist{}, artist_id)
	if artist == nil {
		return db_tables.Artist{}, gorm.ErrRecordNotFound
	}
	return *artist.(*db_tables.Artist), nil
}

// GetArtists returns a list of artists from the database
// Selects the artists with the given artist_ids
func GetArtists(db *gorm.DB, artist_ids []int) ([]db_tables.Artist, error) {
	records := db_model.GetRecordsFromIds(db, &db_tables.Artist{}, artist_ids)
	if records == nil {
		return nil, gorm.ErrRecordNotFound
	}
	artists := make([]db_tables.Artist, len(records))
	for i, record := range records {
		artists[i] = record.(db_tables.Artist)
	}

	return artists, nil
}

func GetArtistsFromPartialPseudo(db *gorm.DB, filters map[string]interface{}, pseudo string) ([]db_tables.Artist, error) {
	records := db_model.GetRecordsByFieldsWithCondition(db, &db_tables.Artist{}, filters, "pseudo LIKE ?", "%"+pseudo+"%")
	if records == nil {
		return nil, gorm.ErrRecordNotFound
	}
	artists := make([]db_tables.Artist, len(records))
	for i, record := range records {
		artists[i] = record.(db_tables.Artist)
	}

	return artists, nil
}
