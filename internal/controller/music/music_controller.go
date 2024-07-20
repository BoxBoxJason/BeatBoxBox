/*
package music_controller is the controller for the musics.

Contains the logic for the music handling. Handles the connection between the API and the database.
*/

package music_controller

import (
	db_tables "BeatBoxBox/internal/model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"encoding/json"
	"path/filepath"
)

// Create the music directory & musics illustrations directory if it doesn't exist
func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "musics"))
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "musics"))
}

// ConvertMusicToJSON converts a music to a JSON object
func ConvertMusicToJSON(music *db_tables.Music) ([]byte, error) {
	artists_ids := make([]int, len(music.Artists))
	for i, artist := range music.Artists {
		artists_ids[i] = artist.Id
	}
	music_json := map[string]interface{}{
		"id":           music.Id,
		"title":        music.Title,
		"lyrics":       music.Lyrics,
		"artists_ids":  artists_ids,
		"album_id":     music.AlbumId,
		"genres":       music.Genres,
		"nb_plays":     music.Nblistened,
		"path":         filepath.Base(music.Path),
		"illustration": filepath.Base(music.Illustration),
		"likes":        music.Likes,
		"created_on":   music.CreatedOn,
		"modified_on":  music.ModifiedOn,
	}
	return json.Marshal(music_json)
}

// ConvertMusicsToJSON converts a slice of musics to a JSON object
func ConvertMusicsToJSON(musics []*db_tables.Music) ([]byte, error) {
	musics_json := make([]map[string]interface{}, len(musics))
	for i, music := range musics {
		artists_ids := make([]int, len(music.Artists))
		for j, artist := range music.Artists {
			artists_ids[j] = artist.Id
		}
		musics_json[i] = map[string]interface{}{
			"id":           music.Id,
			"title":        music.Title,
			"lyrics":       music.Lyrics,
			"artists_ids":  artists_ids,
			"album_id":     music.AlbumId,
			"genres":       music.Genres,
			"nb_plays":     music.Nblistened,
			"path":         filepath.Base(music.Path),
			"illustration": filepath.Base(music.Illustration),
			"likes":        music.Likes,
			"created_on":   music.CreatedOn,
			"modified_on":  music.ModifiedOn,
		}
	}
	return json.Marshal(musics_json)
}
