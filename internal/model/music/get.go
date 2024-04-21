package music_model

import db_controller "BeatBoxBox/internal/controller/db"

// GET METHODS

// GetMusicsFromFilters returns a list of musics from the database
// Filters can be passed to filter the musics
func GetMusicsFromFilters(filters map[string]interface{}) ([]db_controller.Music, error) { // TODO
	return nil, nil
}

// GetMusic returns a music from the database
// Selects the music with the given music_id
func GetMusic(music_id int) (db_controller.Music, error) { // TODO
	return db_controller.Music{}, nil
}

// GetMusics returns a list of musics from the database
// Selects the musics with the given music_ids
func GetMusics(music_ids []int) ([]db_controller.Music, error) { // TODO
	return nil, nil
}
