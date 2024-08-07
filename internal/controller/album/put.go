package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	album_model "BeatBoxBox/internal/model/album"
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
)

// UpdateAlbum updates an album by its id with the given map
func UpdateAlbum(album_id int, album_map map[string]interface{}) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return []byte{}, err
	}
	err = album_model.UpdateAlbum(db, &album, album_map)
	if err != nil {
		return []byte{}, err
	}
	return ConvertAlbumToJSON(&album)
}

func UpdateAlbumArtists(album_id int, artists_ids []int, action string) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	raw_artists, err := artist_model.GetArtists(db, artists_ids)
	if err != nil {
		return []byte{}, err
	} else if len(raw_artists) != len(artists_ids) {
		return []byte{}, custom_errors.NewNotFoundError("Some artists were not found")
	}
	album, err := album_model.GetAlbum(db, album_id)
	if err != nil {
		return []byte{}, err
	}
	artists := make([]*db_tables.Artist, len(raw_artists))
	for i, raw_artist := range raw_artists {
		artists[i] = &raw_artist
	}
	if action == "add" {
		err = album_model.AddArtistsToAlbum(db, &album, artists)
	} else if action == "remove" {
		err = album_model.RemoveArtistsFromAlbum(db, &album, artists)
	} else {
		return []byte{}, custom_errors.NewBadRequestError("Invalid action: " + action)
	}
	if err != nil {
		return []byte{}, err
	}
	return ConvertAlbumToJSON(&album)
}
