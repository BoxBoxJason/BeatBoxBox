package playlist_controller

import (
	db_model "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
)

// PlaylistExists returns whether a playlist exists in the database
func PlaylistExists(playlist_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = playlist_model.GetPlaylist(db, playlist_id)
	return err == nil
}

// PlaylistsExist returns whether a list of playlists exists in the database
func PlaylistsExist(playlist_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylists(db, playlist_ids)
	return err == nil && len(playlists) == len(playlist_ids)
}

// PlaylistExistsFromParams returns whether a playlist exists in the database
func PlaylistExistsFromParams(playlist_name string, playlist_creator_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylistsFromFilters(db, map[string]interface{}{
		"Title":     playlist_name,
		"CreatorId": playlist_creator_id,
	})
	return err == nil && len(playlists) > 0
}

// GetPlaylist returns a playlist from the database
// Selects the playlist with the given playlist_id
func GetPlaylist(playlist_id int) (playlist_model.Playlist, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return playlist_model.Playlist{}, err
	}
	defer db_model.CloseDB(db)
	return playlist_model.GetPlaylist(db, playlist_id)
}

// GetPlaylists returns a list of playlists from the database
// Selects the playlists with the given playlist_ids
func GetPlaylists(playlist_ids []int) ([]playlist_model.Playlist, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	return playlist_model.GetPlaylists(db, playlist_ids)
}
