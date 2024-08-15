package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	playlist_model "BeatBoxBox/internal/model/playlist"
	user_model "BeatBoxBox/internal/model/user"
	db_model "BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

func UpdatePlaylist(playlist_id int, playlist_map map[string]interface{}) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return []byte{}, err
	}
	err = playlist_model.UpdatePlaylist(db, &playlist, playlist_map)
	if err != nil {
		return []byte{}, err
	}
	return ConvertPlaylistToJSON(&playlist)
}

func AddMusicsToPlaylist(playlist_id int, music_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return []byte{}, err
	}
	musics, err := music_model.GetMusics(db, music_ids)
	if err != nil {
		return []byte{}, err
	} else if musics == nil || len(musics) != len(music_ids) {
		return []byte{}, httputils.NewNotFoundError("not all musics found")
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	err = playlist_model.AddMusicsToPlaylist(db, &playlist, musics_ptr)
	if err != nil {
		return []byte{}, err
	}
	return ConvertPlaylistToJSON(&playlist)
}

func RemoveMusicsFromPlaylist(playlist_id int, music_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db.Preload("Musics"), playlist_id)
	if err != nil {
		return []byte{}, err
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
		return []byte{}, httputils.NewNotFoundError("not all musics found in playlist")
	}
	err = playlist_model.RemoveMusicsFromPlaylist(db, &playlist, musics_ptr)
	if err != nil {
		return []byte{}, err
	}
	return ConvertPlaylistToJSON(&playlist)
}

func AddOwnersToPlaylist(playlist_id int, owners_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return []byte{}, err
	}
	new_owners, err := user_model.GetUsers(db, owners_ids)
	if err != nil {
		return []byte{}, err
	}
	new_owners_ptr := make([]*db_tables.User, len(new_owners))
	for i, owner := range new_owners {
		new_owners_ptr[i] = &owner
	}
	err = playlist_model.AddOwnersToPlaylist(db, &playlist, new_owners_ptr)
	if err != nil {
		return []byte{}, err
	}
	return ConvertPlaylistToJSON(&playlist)
}

func RemoveOwnersFromPlaylist(playlist_id int, owners_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return []byte{}, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db.Preload("Owners"), playlist_id)
	if err != nil {
		return []byte{}, err
	}
	owners_ptr := []*db_tables.User{}
	owners_ids_map := map[int]bool{}
	for _, owner_id := range owners_ids {
		owners_ids_map[owner_id] = true
	}
	for _, owner := range playlist.Owners {
		if _, ok := owners_ids_map[owner.Id]; ok {
			owners_ptr = append(owners_ptr, &owner)
		}
	}
	if len(owners_ptr) != len(owners_ids) {
		return []byte{}, httputils.NewNotFoundError("not all owners found in playlist")
	}
	err = playlist_model.RemoveOwnersFromPlaylist(db, &playlist, owners_ptr)
	if err != nil {
		return []byte{}, err
	}
	return ConvertPlaylistToJSON(&playlist)
}
