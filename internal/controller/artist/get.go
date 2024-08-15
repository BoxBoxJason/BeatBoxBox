package artist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
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
	artists := artist_model.GetArtistsFromFilters(db, []string{pseudo}, nil, nil, nil, nil, nil, nil)
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

func GetArtistsJSONFromFilters(pseudos []string, partial_pseudos []string, genres []string, albums_ids []int, albums []string, musics_ids []int, musics []string) ([]byte, error) {
	if len(pseudos)*len(partial_pseudos) != 0 {
		return []byte{}, httputils.NewBadRequestError("Can't use pseudo and partial_pseudo at the same time")
	} else if len(albums_ids)*len(albums) != 0 {
		return []byte{}, httputils.NewBadRequestError("Can't use album_id and album at the same time")
	} else if len(musics_ids)*len(musics) != 0 {
		return []byte{}, httputils.NewBadRequestError("Can't use music_id and music at the same time")
	}
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	artists := artist_model.GetArtistsFromFilters(db.Preload("Musics").Preload("Albums"), pseudos, partial_pseudos, genres, albums_ids, albums, musics_ids, musics)
	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	return ConvertArtistsToJSON(artists_ptr)
}
