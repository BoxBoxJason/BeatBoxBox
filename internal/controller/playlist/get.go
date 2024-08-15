package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	playlist_model "BeatBoxBox/internal/model/playlist"
	db_model "BeatBoxBox/pkg/db_model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	httputils "BeatBoxBox/pkg/utils/httputils"
	"fmt"
	"net/http"
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
	playlist, err := playlist_model.GetPlaylist(db.Preload("Musics").Preload("Subscribers").Preload("Owners"), playlist_id)
	if err != nil {
		return nil, httputils.NewNotFoundError(fmt.Sprintf("Playlist id %d not found", playlist_id))
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
	playlists, err := playlist_model.GetPlaylists(db.Preload("Musics").Preload("Subscribers").Preload("Owners"), playlists_ids)
	if err != nil {
		return nil, err
	} else if len(playlists) == 0 {
		return nil, httputils.NewNotFoundError("some playlists were not found")
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists))
	for i, playlist := range playlists {
		playlists_ptr[i] = &playlist
	}
	return ConvertPlaylistsToJSON(playlists_ptr)
}

func ServePlaylistsFiles(w http.ResponseWriter, playlists_ids []int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	playlists, err := playlist_model.GetPlaylists(db.Preload("Musics"), playlists_ids)
	if err != nil {
		return err
	}
	paths := make(map[string][]string, len(playlists))
	for _, playlist := range playlists {
		paths[playlist.Title] = make([]string, len(playlist.Musics))
		for j, music := range playlist.Musics {
			paths[playlist.Title][j] = file_utils.GetAbsoluteMusicPath(music.Path)
		}
	}
	httputils.ServeSubdirsZip(w, paths, "playlists.zip")
	return nil
}

func ServePlaylistFiles(w http.ResponseWriter, playlist_id int) error {
	db, err := db_model.OpenDB()
	if err != nil {
		return err
	}
	playlist, err := playlist_model.GetPlaylist(db.Preload("Musics"), playlist_id)
	if err != nil {
		return err
	}
	paths := make([]string, len(playlist.Musics))
	for j, music := range playlist.Musics {
		paths[j] = file_utils.GetAbsoluteMusicPath(music.Path)
	}
	httputils.ServeZip(w, paths, playlist.Title+".zip")
	return nil
}

func GetPlaylistsJSONByFilters(titles []string, musics []string, owners []string, music_ids []int, owner_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	playlists, err := playlist_model.GetPlaylistsByFilters(db.Preload("Musics").Preload("Subscribers").Preload("Owners"), titles, musics, owners, music_ids, owner_ids)
	if err != nil {
		return nil, err
	}
	playlists_ptr := make([]*db_tables.Playlist, len(playlists))
	for i, playlist := range playlists {
		playlists_ptr[i] = &playlist
	}
	return ConvertPlaylistsToJSON(playlists_ptr)
}
