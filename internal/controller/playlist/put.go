package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	playlist_model "BeatBoxBox/internal/model/playlist"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
)

func UpdatePlaylist(playlist_id int, playlist_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return err
	}
	return playlist_model.UpdatePlaylist(db, &playlist, playlist_map)
}

func AddMusicsToPlaylist(playlist_id int, music_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return err
	}
	musics, err := music_model.GetMusics(db, music_ids)
	if err != nil {
		return err
	} else if musics == nil || len(musics) != len(music_ids) {
		return custom_errors.NewNotFoundError("not all musics found")
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	return playlist_model.AddMusicsToPlaylist(db, &playlist, musics_ptr)
}

func RemoveMusicsFromPlaylist(playlist_id int, music_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db.Preload("Musics"), playlist_id)
	if err != nil {
		return err
	}
	musics_ptr := []*db_tables.Music{}
	musics_ids_map := map[int]bool{}
	for _, music_id := range music_ids {
		musics_ids_map[music_id] = true
	}
	for _, music := range playlist.Musics {
		if _, ok := musics_ids_map[music.Id]; ok {
			musics_ptr = append(musics_ptr, &music)
		}
	}
	if len(musics_ptr) != len(music_ids) {
		return custom_errors.NewNotFoundError("not all musics found in playlist")
	}
	return playlist_model.RemoveMusicsFromPlaylist(db, &playlist, musics_ptr)
}
