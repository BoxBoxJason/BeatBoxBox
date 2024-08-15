package user_controller

import (
	db_tables "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	playlist_model "BeatBoxBox/internal/model/playlist"
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

func UpdateUser(user_id int, user_map map[string]interface{}) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return err
	}
	return user_model.UpdateUser(db, &user, user_map)
}

func AddMusicsToLikedMusics(user_id int, musics_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return err
	}
	musics, err := music_model.GetMusics(db, musics_ids)
	if err != nil {
		return err
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	return user_model.AddLikedMusicsToUser(db, &user, musics_ptr)
}

func RemoveMusicsFromLikedMusics(user_id int, musics_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db.Preload("LikedMusics"), user_id)
	if err != nil {
		return err
	}
	musics_ids_map := make(map[int]bool, len(user.LikedMusics))
	for _, music := range user.LikedMusics {
		musics_ids_map[music.Id] = true
	}
	musics_ptr := make([]*db_tables.Music, len(musics_ids))
	for i, music := range user.LikedMusics {
		if _, ok := musics_ids_map[music.Id]; ok {
			musics_ptr[i] = &music
		}
	}
	if len(musics_ptr) != len(musics_ids) {
		return httputils.NewNotFoundError("some musics were not found in the liked musics")
	}
	return user_model.RemoveLikedMusicsFromUser(db, &user, musics_ptr)
}

func AddPlaylistsToSubscribedPlaylists(user_id int, playlists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return err
	}
	playlists, err := playlist_model.GetPlaylists(db, playlists_ids)
	if err != nil {
		return err
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists))
	for i, playlist := range playlists {
		playlists_ptr[i] = &playlist
	}
	return user_model.AddSubscribedPlaylistsToUser(db, &user, playlists_ptr)
}

func RemovePlaylistsFromSubscribedPlaylists(user_id int, playlists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db.Preload("SubscribedPlaylists"), user_id)
	if err != nil {
		return err
	}
	playlists_ids_map := make(map[int]bool, len(user.SubscribedPlaylists))
	for _, playlist := range user.SubscribedPlaylists {
		playlists_ids_map[playlist.Id] = true
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists_ids))
	for i, playlist := range user.SubscribedPlaylists {
		if _, ok := playlists_ids_map[playlist.Id]; ok {
			playlists_ptr[i] = &playlist
		}
	}
	if len(playlists_ptr) != len(playlists_ids) {
		return httputils.NewNotFoundError("some playlists were not found in the subscribed playlists")
	}
	return user_model.RemoveSubscribedPlaylistsFromUser(db, &user, playlists_ptr)
}

func AddPlaylistsToOwnedPlaylists(user_id int, playlists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db, user_id)
	if err != nil {
		return err
	}
	playlists, err := playlist_model.GetPlaylists(db, playlists_ids)
	if err != nil {
		return err
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists))
	for i, playlist := range playlists {
		playlists_ptr[i] = &playlist
	}
	return user_model.AddOwnedPlaylistsToUser(db, &user, playlists_ptr)
}

func RemovePlaylistsFromOwnedPlaylists(user_id int, playlists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	user, err := user_model.GetUser(db.Preload("Playlists"), user_id)
	if err != nil {
		return err
	}
	playlists_ids_map := make(map[int]bool, len(user.Playlists))
	for _, playlist := range user.Playlists {
		playlists_ids_map[playlist.Id] = true
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists_ids))
	for i, playlist := range user.Playlists {
		if _, ok := playlists_ids_map[playlist.Id]; ok {
			playlists_ptr[i] = &playlist
		}
	}
	if len(playlists_ptr) != len(playlists_ids) {
		return httputils.NewNotFoundError("some playlists were not found in the owned playlists")
	}
	return user_model.RemoveOwnedPlaylistsFromUser(db, &user, playlists_ptr)
}
