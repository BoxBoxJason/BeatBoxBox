package music_controller

import (
	db_tables "BeatBoxBox/internal/model"
	music_model "BeatBoxBox/internal/model/music"
	db_model "BeatBoxBox/pkg/db_model"
	custom_errors "BeatBoxBox/pkg/errors"
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

// MusicsExists checks if musics exist in the database
func MusicsExist(music_ids []int) bool {
	db, err := db_model.OpenDB()
	if err != nil {
		return false
	}
	defer db_model.CloseDB(db)
	musics, err := music_model.GetMusics(db, music_ids)
	return err == nil && len(musics) == len(music_ids)
}

// GetMusic returns a music from the database in JSON format
func GetMusicJSON(music_id int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	music, err := music_model.GetMusic(db.Preload("Artists"), music_id)
	if err != nil {
		return nil, err
	}
	return ConvertMusicToJSON(&music)
}

// GetMusics returns a list of musics from the database in JSON format
func GetMusicsJSON(musics_ids []int) ([]byte, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)
	musics, err := music_model.GetMusics(db.Preload("Artists"), musics_ids)
	if err != nil {
		return nil, err
	}
	musics_ptr := make([]*db_tables.Music, len(musics))
	for i, music := range musics {
		musics_ptr[i] = &music
	}
	return ConvertMusicsToJSON(musics_ptr)
}

// GetMusicPathFromId returns the path of a music from the database
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

// GetMusicsPathFromIds returns the paths of musics from the database
func GetMusicsPathFromIds(music_ids []int) (map[int]string, error) {
	db, err := db_model.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db_model.CloseDB(db)

	musics, err := music_model.GetMusics(db, music_ids)
	if err != nil {
		return nil, err
	} else if musics == nil || len(musics) != len(music_ids) {
		return nil, custom_errors.NewNotFoundError("some musics were not found")
	}
	paths := make(map[int]string, len(musics))
	for _, music := range musics {
		paths[music.Id] = filepath.Join("data", "musics", music.Path)
	}
	return paths, nil
}
