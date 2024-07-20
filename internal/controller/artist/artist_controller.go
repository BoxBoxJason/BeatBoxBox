package artist_controller

import (
	db_tables "BeatBoxBox/internal/model"
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"encoding/json"
	"path/filepath"
)

// Create the albums illustrations directory if it doesn't exist
func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "artists"))
}

func ConvertArtistToJSON(artist *db_tables.Artist) ([]byte, error) {
	albums_ids := make([]int, len(artist.Albums))
	for i, album := range artist.Albums {
		albums_ids[i] = album.Id
	}
	musics_ids := make([]int, len(artist.Musics))
	for i, music := range artist.Musics {
		musics_ids[i] = music.Id
	}

	artist_json := map[string]interface{}{
		"id":           artist.Id,
		"pseudo":       artist.Pseudo,
		"bio":          artist.Bio,
		"illustration": filepath.Base(artist.Illustration),
		"albums_ids":   albums_ids,
		"musics_ids":   musics_ids,
		"created_on":   artist.CreatedOn,
		"modified_on":  artist.ModifiedOn,
	}
	return json.Marshal(artist_json)
}

func ConvertArtistsToJSON(artists []*db_tables.Artist) ([]byte, error) {
	artists_json := make([]map[string]interface{}, len(artists))
	for i, artist := range artists {
		artist_json, err := ConvertArtistToJSON(artist)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(artist_json, &artists_json[i])
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(artists_json)
}
