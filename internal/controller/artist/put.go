package artist_controller

import (
	db_model "BeatBoxBox/internal/model"
	artist_model "BeatBoxBox/internal/model/artist"
)

func UpdateArtist(artist_id int, artist_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return artist_model.UpdateArtist(db, artist_id, artist_map)
}
