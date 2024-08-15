package artist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"
)

// DeleteArtist deletes an artist from the database
func DeleteArtist(artist_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	artist, err := artist_model.GetArtist(db, artist_id)
	if err != nil {
		return httputils.NewNotFoundError(fmt.Sprintf("artist with id %d not found", artist_id))
	}
	return artist_model.DeleteArtistFromRecord(db, &artist)
}

// DeleteArtists deletes a list of artists from the database
// Selects the artists with the given artist_ids
func DeleteArtists(artist_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	artists, err := artist_model.GetArtists(db, artist_ids)
	if err != nil {
		return httputils.NewNotFoundError("no artist found")
	} else if len(artists) != len(artist_ids) {
		return httputils.NewNotFoundError("some artists were not found")
	}

	artists_ptr := make([]*db_tables.Artist, len(artists))
	for i, artist := range artists {
		artists_ptr[i] = &artist
	}
	return artist_model.DeleteArtistsFromRecords(db, artists_ptr)
}
