package artist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
)

// ArtistExists checks if an artist with the given artist_id exists
func ArtistExists(artist_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = artist_model.GetArtist(db, artist_id)
	return err == nil
}

// ArtistsExist checks if all the artists with the given artist_ids exist
func ArtistsExist(artist_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	artists, err := artist_model.GetArtists(db, artist_ids)
	return err == nil && len(artists) == len(artist_ids)
}

// IsPseudoTaken checks if a pseudo is already taken by an artist
func IsPseudoTaken(pseudo string) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	artists := artist_model.GetArtistsFromFilters(db, map[string]interface{}{"pseudo": pseudo})
	return artists == nil && len(artists) > 0
}

// GetArtist returns an artist from the database in JSON format
func GetArtistJSON(artist_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	music, err := artist_model.GetArtist(db.Preload("Musics").Preload("Albums"), artist_id)
	if err != nil {
		return nil, err
	}
	return ConvertArtistToJSON(&music)
}

// GetArtists returns a list of artists from the database in JSON format
func GetArtistsJSON(artists_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	artists, err := artist_model.GetArtists(db.Preload("Musics").Preload("Albums"), artists_ids)
	if err != nil {
		return nil, err
	}
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	return ConvertArtistsToJSON(artists_ptr)
}

func GetArtistsJSONFromFilters(filters map[string]interface{}) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	artists := artist_model.GetArtistsFromFilters(db.Preload("Musics").Preload("Albums"), filters)
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	return ConvertArtistsToJSON(artists_ptr)
}
