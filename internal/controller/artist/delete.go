package artist_controller

import "errors"

// DeleteArtist deletes an artist from the database
// Selects the artist with the given artist_id
func DeleteArtist(artist_id int) error {
	if !ArtistExists(artist_id) {
		return errors.New("artist does not exist")
	}
	return DeleteArtist(artist_id)
}

// DeleteArtists deletes a list of artists from the database
// Selects the artists with the given artist_ids
func DeleteArtists(artist_ids []int) error {
	if !ArtistsExist(artist_ids) {
		return errors.New("at least one artist does not exist")
	}
	return DeleteArtists(artist_ids)
}
