package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	"BeatBoxBox/pkg/db_model"
	httputils "BeatBoxBox/pkg/utils/httputils"
)

// DeletePlaylist deletes a playlist from the database
func DeletePlaylist(playlist_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return err
	}
	return playlist_model.DeletePlaylist(db, &playlist)
}

// DeletePlaylists deletes a list of playlists from the database
func DeletePlaylists(playlist_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylists(db, playlist_ids)
	if err != nil {
		return err
	} else if playlists == nil || len(playlists) != len(playlist_ids) {
		return httputils.NewNotFoundError("not all playlists found")
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists))
	for i, playlist := range playlists {
		playlists_ptr[i] = &playlist
	}
	return playlist_model.DeletePlaylists(db, playlists_ptr)
}
