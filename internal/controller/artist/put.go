package artist_controller

import (
	artist_model "BeatBoxBox/internal/model/artist"
	"BeatBoxBox/pkg/db_model"
)

func UpdateArtist(artist_id int, artist_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	artist, err := artist_model.GetArtist(db, artist_id)
	if err != nil {
		return err
	}
	return artist_model.UpdateArtist(db, &artist, artist_map)
}
