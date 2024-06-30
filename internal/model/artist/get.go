package artist_model

import (
	db_tables "BeatBoxBox/internal/model"
	"BeatBoxBox/pkg/db_model"
	"gorm.io/gorm"
)

// GetArtistsFromFilters returns a list of artists from the database
// Filters can be passed to filter the artists
func GetArtistsFromFilters(db *gorm.DB, filters map[string]interface{}) []db_tables.Artist {
	raw_artists := db_model.GetRecordsByFields(db, &db_tables.Artist{}, filters)
	if raw_artists == nil {
		return nil
	}
	artists := make([]db_tables.Artist, len(raw_artists))
	for i, artist := range raw_artists {
		artists[i] = artist.(db_tables.Artist)
	}

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
