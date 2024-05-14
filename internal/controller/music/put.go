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

func UpdateMusicGenres(music_id int, genres []string) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return music_model.UpdateMusic(db, music_id, map[string]interface{}{"genres": genres})
}

func AddArtistsToMusic(music_id int, artists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return music_model.AddArtistsToMusic(db, music_id, artists_ids)
}

func RemoveArtistFromMusic(music_id int, artists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return music_model.RemoveArtistsFromMusic(db, music_id, artists_ids)
}
