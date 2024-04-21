package music_model

// GET METHODS

// GetMusicsFromFilters returns a list of musics from the database
// Filters can be passed to filter the musics
func GetMusicsFromFilters(filters map[string]interface{}) ([]Music, error) { // TODO
	return nil, nil
}

// GetMusic returns a music from the database
// Selects the music with the given music_id
func GetMusic(music_id int) (Music, error) { // TODO
	return Music{}, nil
}

// GetMusics returns a list of musics from the database
// Selects the musics with the given music_ids
func GetMusics(music_ids []int) ([]Music, error) { // TODO
	return nil, nil
}
