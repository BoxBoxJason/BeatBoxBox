package music_controller

import (
	db_model "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
)

func UpdateMusic(music_id int, music_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return music_model.UpdateMusic(db, music_id, music_map)
}
