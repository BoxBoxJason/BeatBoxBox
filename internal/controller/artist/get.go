package artist_controller

import (
	db_model "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	"encoding/json"
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
	return err != nil && len(artists) == len(artist_ids)
}

// IsPseudoTaken checks if a pseudo is already taken by an artist
func IsPseudoTaken(pseudo string) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = artist_model.GetArtistByPseudo(db, pseudo)
	return err == nil
}

// GetArtist returns an artist from the database
// Selects the artist with the given artist_id
// Returns the artist as a JSON object
func GetArtist(artist_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	music, err := artist_model.GetArtist(db, artist_id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(music)
}

// GetArtists returns a list of artists from the database
// Selects the artists with the given artist_ids
// Returns the artists as a JSON array
func GetArtists(artists_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	artists, err := artist_model.GetArtists(db, artists_ids)
	if err != nil {
		return nil, err
	}
	return json.Marshal(artists)
}
