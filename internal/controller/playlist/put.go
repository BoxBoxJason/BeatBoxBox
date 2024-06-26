package playlist_controller

import (
	db_model "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
)

func UpdatePlaylist(playlist_id int, playlist_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return playlist_model.UpdatePlaylist(db, playlist_id, playlist_map)
}

func AddMusicsToPlaylist(playlist_id int, music_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return playlist_model.AddMusicsToPlaylist(db, playlist_id, music_ids)
}

func RemoveMusicsFromPlaylist(playlist_id int, music_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	return playlist_model.RemoveMusicsFromPlaylist(db, playlist_id, music_ids)
}
