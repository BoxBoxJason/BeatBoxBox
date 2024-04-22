package music_model

import (
	"gorm.io/gorm"
)

// GET METHODS

// GetMusicsFromFilters returns a list of musics from the database
// Filters can be passed to filter the musics
func GetMusicsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]Music, error) {
	var musics []Music
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&musics)
	if result.Error != nil {
		return nil, result.Error
	}
	return musics, nil
}

// GetMusic returns a music from the database
// Selects the music with the given music_id
func GetMusic(db *gorm.DB, music_id int) (Music, error) {
	var music Music
	// Using `First` to retrieve the first record that matches the music_id
	result := db.Where("Id = ?", music_id).First(&music)
	if result.Error != nil {
		return Music{}, result.Error
	}
	return music, nil
}

// GetMusics returns a list of musics from the database
// Selects the musics with the given music_ids
func GetMusics(db *gorm.DB, music_ids []int) ([]Music, error) {
	var musics []Music
	// Using `Find` to retrieve records with the IDs in music_ids slice
	result := db.Where("Id IN ?", music_ids).Find(&musics)
	if result.Error != nil {
		return nil, result.Error
	}
	return musics, nil
}
