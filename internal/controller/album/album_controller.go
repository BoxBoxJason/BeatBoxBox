package album_controller

import (
	db_tables "BeatBoxBox/internal/model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"encoding/json"
	"path/filepath"
)

// Create the illustrations directory if it doesn't exist
func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "albums"))
}

// ConvertAlbumToJSON converts an album to a JSON object
func ConvertAlbumToJSON(album *db_tables.Album) ([]byte, error) {
	musics_ids := make([]int, len(album.Musics))
	for i, music := range album.Musics {
		musics_ids[i] = music.Id
	}
	artists_ids := make([]int, len(album.Artists))
	for i, artist := range album.Artists {
		artists_ids[i] = artist.Id
	}

	album_json := map[string]interface{}{
		"id":           album.Id,
		"title":        album.Title,
		"description":  album.Description,
		"illustration": filepath.Base(album.Illustration),
		"artists_ids":  artists_ids,
		"musics_ids":   musics_ids,
		"created_on":   album.CreatedOn,
		"modified_on":  album.ModifiedOn,
	}
	return json.Marshal(album_json)
}

// ConvertAlbumsToJSON converts a list of albums to a JSON array
func ConvertAlbumsToJSON(albums []*db_tables.Album) ([]byte, error) {
	albums_json := make([]map[string]interface{}, len(albums))
	for i, album := range albums {
		musics_ids := make([]int, len(album.Musics))
		for i, music := range album.Musics {
			musics_ids[i] = music.Id
		}
		artists_ids := make([]int, len(album.Artists))
		for i, artist := range album.Artists {
			artists_ids[i] = artist.Id
		}
		albums_json[i] = map[string]interface{}{
			"id":           album.Id,
			"title":        album.Title,
			"description":  album.Description,
			"illustration": filepath.Base(album.Illustration),
			"artists_ids":  artists_ids,
			"musics_ids":   musics_ids,
			"created_on":   album.CreatedOn,
			"modified_on":  album.ModifiedOn,
		}
	}
	return json.Marshal(albums_json)
}
