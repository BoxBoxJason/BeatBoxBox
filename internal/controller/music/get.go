package music_controller

import (
	db_model "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	"encoding/json"
	"path/filepath"
)

// MusicExists checks if a music exists in the database
func MusicExists(music_id int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	_, err = music_model.GetMusic(db, music_id)
	return err == nil
}

// GetMusic returns a music from the database
// Selects the music with the given music_id
// Returns the music as a JSON object
func GetMusic(music_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(music)
}

// GetMusics returns a list of musics from the database
// Selects the musics with the given music_ids
// Returns the musics as a JSON array
func GetMusics(musics_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	musics, err := music_model.GetMusics(db, musics_ids)
	if err != nil {
		return nil, err
	}
	return json.Marshal(musics)
}

func GetMusicPathFromId(music_id int) (string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return "", err
	}
	defer db_model.CloseDB(db)

	music, err := music_model.GetMusic(db, music_id)
	if err != nil {
		return "", err
	}
	return filepath.Join("data", "musics", music.Path), nil
}

func GetMusicsPathFromIds(music_ids []int) ([]string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)

	musics, err := music_model.GetMusics(db, music_ids)
	if err != nil {
		return nil, err
	}
	paths := []string{}
	for _, music := range musics {
		paths = append(paths, filepath.Join("data", "musics", music.Path))
	}
	return paths, nil
}
