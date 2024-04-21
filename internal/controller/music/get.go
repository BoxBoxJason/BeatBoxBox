package music_controller

import (
	music_model "BeatBoxBox/internal/model/music"
	"encoding/json"
	"path/filepath"
)

func musicExists(title string) bool {
	return false
}

// GetMusic returns a music from the database
// Selects the music with the given music_id
// Returns the music as a JSON object
func GetMusic(music_id int) ([]byte, error) {
	music, err := music_model.GetMusic(music_id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(music)
}

// GetMusics returns a list of musics from the database
// Selects the musics with the given music_ids
// Returns the musics as a JSON array
func GetMusics(musics_ids []int) ([]byte, error) {
	musics, err := music_model.GetMusics(musics_ids)
	if err != nil {
		return nil, err
	}
	return json.Marshal(musics)
}

func GetMusicPathFromId(music_id int) (string, error) {
	music, err := music_model.GetMusic(music_id)
	if err != nil {
		return "", err
	}
	return music.Path, nil
}

func GetMusicsPathFromIds(music_ids []int) ([]string, error) {
	musics, err = music_model.GetMusics(music_ids)
	if err != nil {
		return nil, err
	}
	paths := []string{}
	for _, music := range musics {
		paths = append(paths, filepath.Join("data", "musics", music.Path))
	}
	return paths, nil
}
