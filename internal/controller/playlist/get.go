package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
	"fmt"
	"path"
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

// GetPlaylist returns a playlist from the database as a JSON object
func GetPlaylistJSON(playlist_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db.Preload("Musics"), playlist_id)
	if err != nil {
		return nil, custom_errors.NewNotFoundError(fmt.Sprintf("Playlist id %d not found", playlist_id))
	}
	return ConvertPlaylistToJSON(&playlist)
}

// GetPlaylists returns a list of playlists from the database as a JSON array
func GetPlaylistsJSON(playlists_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylists(db.Preload("Musics"), playlists_ids)
	if err != nil {
		return nil, err
	} else if playlists == nil || len(playlists) == 0 {
		return nil, custom_errors.NewNotFoundError("some playlists were not found")
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists))
	for i, playlist := range playlists {
		playlists_ptr[i] = &playlist
	}
	return ConvertPlaylistsToJSON(playlists_ptr)
}

func GetMusicsPathFromPlaylist(playlist_id int) (string, []string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return "", nil, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db.Preload("Musics"), playlist_id)
	if err != nil {
		return "", nil, err
	}
	musics_paths := []string{}
	for _, music := range playlist.Musics {
		musics_paths = append(musics_paths, path.Join("data", "musics", music.Path))
	}
	return playlist.Title, musics_paths, nil
}

func GetMusicsPathFromPlaylists(playlist_ids []int) (map[string][]string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylists(db.Preload("Musics"), playlist_ids)
	if err != nil {
		return nil, err
	}
	musics_paths := make(map[string][]string, len(playlists))
	for _, playlist := range playlists {
		musics_paths[playlist.Title] = make([]string, len(playlist.Musics))
		for i, music := range playlist.Musics {
			musics_paths[playlist.Title][i] = path.Join("data", "musics", music.Path)
		}
	}
	return musics_paths, nil
}
