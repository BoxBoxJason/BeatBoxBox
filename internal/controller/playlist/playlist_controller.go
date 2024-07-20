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
	playlist_json := map[string]interface{}{
		"id":           playlist.Id,
		"title":        playlist.Title,
		"description":  playlist.Description,
		"illustration": playlist.Illustration,
		"creator_id":   playlist.CreatorId,
		"musics_ids":   musics_ids,
		"protected":    playlist.Protected,
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
		playlists_json[i] = map[string]interface{}{
			"id":           playlist.Id,
			"title":        playlist.Title,
			"description":  playlist.Description,
			"illustration": playlist.Illustration,
			"creator_id":   playlist.CreatorId,
			"musics_ids":   musics_ids,
			"protected":    playlist.Protected,
			"created_on":   playlist.CreatedOn,
			"modified_on":  playlist.ModifiedOn,
		}
	}
	return json.Marshal(playlists_json)
}
