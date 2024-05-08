package playlist_controller

import (
	db_model "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	custom_errors "BeatBoxBox/pkg/errors"
	"encoding/json"
	"path"
	"strconv"
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

// GetPlaylist returns a playlist from the database as a JSON object
// Selects the playlist with the given playlist_id
func GetPlaylist(playlist_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return nil, custom_errors.NewNotFoundError("Playlist id " + strconv.Itoa(playlist_id) + " not found")
	}
	return json.Marshal(playlist)
}

// GetPlaylists returns a list of playlists from the database as a JSON array
// Selects the playlists with the given playlist_ids
func GetPlaylists(playlist_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylists(db, playlist_ids)
	if err != nil {
		return nil, custom_errors.NewNotFoundError("Playlists not found")
	}
	return json.Marshal(playlists)
}

func GetMusicsPathFromPlaylist(playlist_id int) (string, []string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return "", nil, err
	}
	defer db_model.CloseDB(db)
	playlist, err := playlist_model.GetPlaylist(db, playlist_id)
	if err != nil {
		return "", nil, custom_errors.NewNotFoundError("Playlist id " + strconv.Itoa(playlist_id) + " not found")
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
	playlists, err := playlist_model.GetPlaylists(db, playlist_ids)
	if err != nil {
		return nil, err
	}
	musics_paths := map[string][]string{}
	for _, playlist := range playlists {
		musics_paths[playlist.Title] = []string{}
		for _, music := range playlist.Musics {
			musics_paths[playlist.Title] = append(musics_paths[playlist.Title], path.Join("data", "musics", music.Path))
		}
	}
	return musics_paths, nil
}
