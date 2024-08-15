package playlist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"encoding/json"
	"path/filepath"
)

// Create the music directory if it doesn't exist
func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "playlists"))
}

func ConvertPlaylistToJSON(playlist *db_tables.Playlist) ([]byte, error) {
	musics_ids := make([]int, len(playlist.Musics))
	for i, music := range playlist.Musics {
		musics_ids[i] = music.Id
	}
	owners_ids := make([]int, len(playlist.Owners))
	for i, owner := range playlist.Owners {
		owners_ids[i] = owner.Id
	}
	subscribers_ids := make([]int, len(playlist.Subscribers))
	for i, subscriber := range playlist.Subscribers {
		subscribers_ids[i] = subscriber.Id
	}
	playlist_json := map[string]interface{}{
		"id":           playlist.Id,
		"title":        playlist.Title,
		"description":  playlist.Description,
		"illustration": playlist.Illustration,
		"owners_ids":   owners_ids,
		"subscribers":  subscribers_ids,
		"musics_ids":   musics_ids,
		"public":       playlist.Public,
		"created_on":   playlist.CreatedOn,
		"modified_on":  playlist.ModifiedOn,
	}
	return json.Marshal(playlist_json)
}

func ConvertPlaylistsToJSON(playlists []*db_tables.Playlist) ([]byte, error) {
	playlists_json := make([]map[string]interface{}, len(playlists))
	for i, playlist := range playlists {
		musics_ids := make([]int, len(playlist.Musics))
		for j, music := range playlist.Musics {
			musics_ids[j] = music.Id
		}
		owners_ids := make([]int, len(playlist.Owners))
		for j, owner := range playlist.Owners {
			owners_ids[j] = owner.Id
		}
		subscribers_ids := make([]int, len(playlist.Subscribers))
		for j, subscriber := range playlist.Subscribers {
			subscribers_ids[j] = subscriber.Id
		}
		playlists_json[i] = map[string]interface{}{
			"id":           playlist.Id,
			"title":        playlist.Title,
			"description":  playlist.Description,
			"illustration": playlist.Illustration,
			"owners_ids":   owners_ids,
			"subscribers":  subscribers_ids,
			"musics_ids":   musics_ids,
			"public":       playlist.Public,
			"created_on":   playlist.CreatedOn,
			"modified_on":  playlist.ModifiedOn,
		}
	}
	return json.Marshal(playlists_json)
}
