package music_model

import (
	db_model "BeatBoxBox/internal/model"

	"gorm.io/gorm"
)

// GetMusicsFromFilters returns a list of musics from the database
// Filters can be passed to filter the musics
func GetMusicsFromFilters(db *gorm.DB, filters map[string]interface{}) ([]db_model.Music, error) {
	var musics []db_model.Music
	// Apply filters and retrieve the records
	result := db.Where(filters).Find(&musics)
	if result.Error != nil {
		return nil, result.Error
	}
	return musics, nil
}

// GetMusic returns a music from the database
// Selects the music with the given music_id
func GetMusic(db *gorm.DB, music_id int) (db_model.Music, error) {
	var music db_model.Music
	// Using `First` to retrieve the first record that matches the music_id
	result := db.Where("id = ?", music_id).First(&music)
	if result.Error != nil {
		return db_model.Music{}, result.Error
	}
	return music, nil
}

// GetMusics returns a list of musics from the database
// Selects the musics with the given music_ids
func GetMusics(db *gorm.DB, music_ids []int) ([]db_model.Music, error) {
	var musics []db_model.Music
	// Using `Find` to retrieve records with the IDs in music_ids slice
	result := db.Where("id IN ?", music_ids).Find(&musics)
	if result.Error != nil {
		return nil, result.Error
	}
	return musics, nil
}
